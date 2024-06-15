package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, fmt.Errorf("failed to open file, please check if file exists at %v", fm.InputFilePath)
	}

	defer file.Close() // func is called after the ReadLines finishes

	scanner := bufio.NewScanner(file)

	var lines []string

	// would continue looping on every line of the file till
	// there is no line left
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("failed to read file content")
	}

	// file.Close()
	return lines, nil

}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

func New(InputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath,
		outputFilePath,
	}
}
