package filehandler

import (
	"encoding/json"
	"log"
	"os"
	"poller-app/internal/domain"
)

// writes json rows on newline on a given file
func WriteFile(data domain.Response, fileName string, logger *log.Logger) error {
	err := checkFile(fileName, logger)
	if err != nil {
		return err
	}
	readBytes, err := ReadFile(fileName, logger)
	if err != nil {
		log.Println(err)
		return err
	}
	var response []domain.Response
	if len(readBytes) > 0 {
		err = json.Unmarshal(readBytes, &response)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	response = append(response, data)
	writeBytes, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return err
	}
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	n, err := f.Write(writeBytes)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("wrote %d bytes", n)
	return nil
}

// Read the complete file into a buffer
func ReadFile(fileName string, logger *log.Logger) ([]byte, error) {
	err := checkFile(fileName, logger)
	if err != nil {
		return nil, err
	}
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	return fileContents, nil
}

// check if the file exists or not
func checkFile(filename string, logger *log.Logger) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			logger.Println(err)
			return err
		}
	}
	return nil
}
