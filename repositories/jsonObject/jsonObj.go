package jsonobject

import (
	"io/ioutil"
	"log"
	"os"
	"user_api/config"
)

func OpenJsonFile() (*[]byte, error) {
	// get confing
	cfg := config.Get()

	// open file
	jsonFile, err := os.OpenFile(cfg.JsonPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Printf("Open json file: %v", err)
		return nil, err
	}
	defer jsonFile.Close()

	// read bytes:
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Read json file: %v", err)
		return nil, err
	}

	return &bytes, nil
}

// TODO: НЕ НУЖНА!!!!!!!!
func WriteJsonFile(jsonBytes []byte) error {

	if err := ioutil.WriteFile(config.Get().JsonPath, jsonBytes, 0660); err != nil {
		log.Printf("write json file: %v", err)
		return err
	}

	return nil
}
