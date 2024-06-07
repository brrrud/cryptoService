package contollers

import (
	"bytes"
	"cryptoService/cryptography/utils"
	"cryptoService/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

const (
	uploadURL  = "http://example.com/upload" // URL для отправки данных
	setupURL   = "URL"
	numWorkers = 5 // Количество воркеров для чтения файла
)

func LoadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters
	pathForLoadFile := r.URL.Query().Get("pathForLoadFile")
	cryptoAlgorithm := r.URL.Query().Get("cryptoAlgorithm")
	padding := r.URL.Query().Get("padding")
	cipherMode := r.URL.Query().Get("cipherMode")
	content := r.URL.Query().Get("content")
	format := r.URL.Query().Get("format")
	//countParts := r.URL.Query().Get("countParts")
	key := r.URL.Query().Get("key")

	// Call LoadFile with the extracted parameters
	go LoadFile(pathForLoadFile, cryptoAlgorithm, padding, cipherMode, content, format, key)

	// Respond to the client
	fmt.Fprintln(w, "File loading started")
}

func LoadFile(pathForLoadFile string, cryptoAlgorithm, padding, cipherMode, content, format string, key string) {
	numWorkers := 5
	err := SendSetupMessage(cryptoAlgorithm, padding, cipherMode, content, format, numWorkers)
	if err != nil {
		fmt.Println("Error with SendSetupMessage")
		return
	}

	paddingAlgo := utils.GetPaddingByName(padding)
	cipherAlgo := utils.GetAlgorithmByName(cryptoAlgorithm)
	modeAlgo := utils.GetModeByName(cipherMode)

	cryptoSystem, _ := utils.NewCipherModeBuilder().SetMode(modeAlgo).SetPadding(paddingAlgo).SetEncryptionAlgorithm(cipherAlgo).SetKey(key).Build()

	pfr, err := utils.NewParallelFileReader(pathForLoadFile, numWorkers, cryptoSystem.Encrypt)
	if err != nil {
		fmt.Printf("Error creating ParallelFileReader: %v\n", err)
		return
	}
	pfr.Read()

	var wg sync.WaitGroup

	for fileBlocks := range pfr.ResultChan {
		for _, block := range fileBlocks {
			wg.Add(1)
			go func(b utils.FileBlock) {
				defer wg.Done()
				err := SendBlock(b, uploadURL)
				if err != nil {
					fmt.Printf("Error sending block: %v\n", err)
				}
			}(block)
		}
	}
	wg.Wait()
}

func SendBlock(block utils.FileBlock, uploadURL string) error {
	messagePart := models.MessagePart{
		MsgPartID: int(block.Index),
		Data:      string(block.Data),
	}

	jsonData, err := json.Marshal(messagePart)
	if err != nil {
		return fmt.Errorf("failed to marshal block: %w", err)
	}

	// Отправка данных на сервер
	resp, err := http.Post(uploadURL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send block: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %s", resp.Status)
	}
	return nil
}

func SendSetupMessage(cryptoAlgorithm, padding, cipherMode, content, format string, countParts int) error {
	message := models.Message{
		CryptoAlgorithm: cryptoAlgorithm,
		Padding:         padding,
		CipherMode:      cipherMode,
		Content:         content,
		Format:          format,
		CountParts:      countParts,
	}

	// Преобразование структуры Message в JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal setup message: %w", err)
	}

	// Отправка данных на сервер
	resp, err := http.Post(setupURL, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send setup message: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned non-200 status: %s", resp.Status)
	}

	return nil
}

type MessageController struct {
	mu          sync.Mutex
	totalChunks int
	chunks      [][]byte
	received    int
}

func NewMessageController(totalChunks int) *MessageController {
	return &MessageController{
		totalChunks: totalChunks,
		chunks:      make([][]byte, totalChunks),
		received:    0,
	}
}

func (mc *MessageController) AddChunk(w http.ResponseWriter, r *http.Request) {
	chunkIndexStr := r.URL.Query().Get("index")
	if chunkIndexStr == "" {
		http.Error(w, "Missing chunk index", http.StatusBadRequest)
		return
	}

	chunkIndex, err := strconv.Atoi(chunkIndexStr)
	if err != nil || chunkIndex < 0 || chunkIndex >= mc.totalChunks {
		http.Error(w, "Invalid chunk index", http.StatusBadRequest)
		return
	}

	chunk, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	mc.mu.Lock()
	defer mc.mu.Unlock()

	if mc.chunks[chunkIndex] == nil {
		mc.chunks[chunkIndex] = chunk
		mc.received++
		fmt.Fprintf(w, "Chunk %d received, total received: %d/%d", chunkIndex, mc.received, mc.totalChunks)
	} else {
		http.Error(w, "Chunk already received", http.StatusBadRequest)
	}
}

func (mc *MessageController) GetMessage(w http.ResponseWriter, r *http.Request) {
	mc.mu.Lock()
	defer mc.mu.Unlock()

	if mc.received < mc.totalChunks {
		http.Error(w, "Message not fully received", http.StatusBadRequest)
		return
	}

	var fullMessage []byte
	for _, chunk := range mc.chunks {
		fullMessage = append(fullMessage, chunk...)
	}

	w.Write(fullMessage)
}
