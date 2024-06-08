package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

type FileBlock struct {
	Index int64
	Data  []byte
}

type ParallelFileReader struct {
	filePath             string
	numWorkers           int
	file                 *os.File
	fileSize             int64
	dataChan             chan FileBlock
	ResultChan           chan []FileBlock
	errChan              chan error
	wg                   sync.WaitGroup
	processDataChunkFunc func(block []byte) ([]byte, error)
	blockSize            int64
	fileExtension        string
}

func NewParallelFileReader(filePath string, numWorkers int, fnc func(block []byte) ([]byte, error)) (*ParallelFileReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("error getting file info: %v", err)
	}

	return &ParallelFileReader{
		filePath:             filePath,
		numWorkers:           numWorkers,
		file:                 file,
		fileSize:             fileInfo.Size(),
		dataChan:             make(chan FileBlock, numWorkers),
		ResultChan:           make(chan []FileBlock),
		errChan:              make(chan error, 1),
		processDataChunkFunc: fnc,
		blockSize:            fileInfo.Size() / int64(numWorkers),
		fileExtension:        filepath.Ext(filePath),
	}, nil
}

// Read запускает параллельное чтение файла
func (pfr *ParallelFileReader) Read() {
	chunkSize := pfr.fileSize / int64(pfr.numWorkers)

	// Запуск горутин для параллельного чтения файла
	for i := 0; i < pfr.numWorkers; i++ {
		pfr.wg.Add(1)
		go pfr.readChunk(i, chunkSize)
	}
	var mu sync.Mutex
	var results []FileBlock
	go func() {
		for block := range pfr.dataChan {
			processedBlocks, err := pfr.processDataChunkFunc(block.Data)
			if err != nil {
				pfr.errChan <- err
				return
			}
			mu.Lock()
			results = append(results, FileBlock{Index: block.Index, Data: processedBlocks})
			mu.Unlock()
		}
		pfr.ResultChan <- results
		close(pfr.ResultChan)
	}()
	// Ожидание завершения всех горутин
	go func() {
		pfr.wg.Wait()
		close(pfr.dataChan)
		close(pfr.errChan)
	}()
}

// readChunk читает часть файла
func (pfr *ParallelFileReader) readChunk(workerID int, chunkSize int64) {
	defer pfr.wg.Done()
	start := int64(workerID) * chunkSize
	end := start + chunkSize
	if workerID == pfr.numWorkers-1 {
		end = pfr.fileSize // Последний воркер читает до конца файла
	}
	//blockIndex := start / blockSize
	for offset := start; offset < end; offset += pfr.blockSize {
		buf := make([]byte, pfr.blockSize)
		if end-offset < pfr.blockSize {
			buf = make([]byte, end-offset)
		}
		n, err := pfr.file.ReadAt(buf, offset)
		if err != nil && err != io.EOF {
			pfr.errChan <- err
			return
		}
		if n > 0 {
			blockIndex := offset / pfr.blockSize
			pfr.dataChan <- FileBlock{Index: blockIndex, Data: buf[:n]}
			//blockIndex++
		}
		if err == io.EOF {
			break
		}
	}
}

func (pfr *ParallelFileReader) Close() {
	if err := <-pfr.errChan; err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	} else {
		fmt.Println("File reading completed successfully")
	}
	err := pfr.file.Close()
	if err != nil {
		return
	}
}

//func (pfr *ParallelFileReader) BuildFileByChunks(resultChan chan, ) {
//
//}

func ReadFileToBytes(filePath string) ([]byte, error) {
	const maxSize int64 = 2 * 1024 * 1024 // 2 MB

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileInfo.Size() > maxSize {
		return nil, fmt.Errorf("file incorrect size , %d", fileInfo.Size())
	}

	byteArray := make([]byte, fileInfo.Size())

	_, err = io.ReadFull(file, byteArray)
	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func GetFileExtension(filePath string) string {
	return filepath.Ext(filePath)
}
