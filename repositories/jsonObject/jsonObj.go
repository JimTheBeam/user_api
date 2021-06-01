package jsonobject

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"user_api/config"
	"user_api/model"
)

// JsonFile opens json file and unmarshal to struct
func JsonFile() (*model.Users, error) {
	bytes, err := openJsonFile()
	if err != nil {
		log.Printf("open json file: %v", err)
		return &model.Users{}, err
	}
	var users model.Users

	// unmarshal bytes to struct
	if err := json.Unmarshal(*bytes, &users); err != nil {
		log.Printf("JsonFile: unmarshal bytes: %v", err)
		return &model.Users{}, err
	}

	return &users, nil

}

// openJsonFile opens json file
func openJsonFile() (*[]byte, error) {
	// get confing
	cfg := config.Get()

	// open file
	jsonFile, err := os.OpenFile(cfg.JsonPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Printf("Open json file: %v", err)
		return nil, err
	}
	defer jsonFile.Close()

	info, err := jsonFile.Stat()
	if err != nil {
		fmt.Println("EMPTY FILE!!!!!!!!")
		fmt.Println(info)
		return nil, err
	}
	// check if file is empty
	if info.Size() == 0 {
		// create a new empty json object in file
		writeJsonFile(&model.Users{})
	}

	// read bytes:
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Read json file: %v", err)
		return nil, err
	}

	return &bytes, nil
}

// WriteJsonFile write users struct to file
func writeJsonFile(users *model.Users) error {

	jsonBytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Printf("write json file marshal bytes: %v", err)
		return err
	}

	if err := ioutil.WriteFile(config.Get().JsonPath, jsonBytes, 0660); err != nil {
		log.Printf("write json file: %v", err)
		return err
	}

	return nil
}
