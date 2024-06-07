package contollers

import (
	"bytes"
	"cryptoService/cryptography/modes"
	"cryptoService/cryptography/utils"
	"cryptoService/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const (
	uploadURL  = "http://example.com/upload" // URL для отправки данных
	numWorkers = 5                           // Количество воркеров для чтения файла
)

func LoadFile(pathForLoadFile string, mod modes.CipherMode) {
	numWorkers := 5
	pfr, err := utils.NewParallelFileReader(pathForLoadFile, numWorkers, mod.Encrypt)
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
