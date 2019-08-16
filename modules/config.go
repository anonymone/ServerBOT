package modules

import (
	"encoding/json"
	"log"
	"os"
)

// JSONOperator golint ask me to write something
type JSONOperator interface {
	ReadJson(string) (interface{}, error)
	WriteJson(string) error
	test()
}

// ConfigInformation golint asked me to write something here.
type ConfigInformation struct {
	Receiver   []string
	Sender     string
	ServerHost string
	ServerPort int
	APICredit  string
	// Task Schedule
	TaskSend  string
	TaskCheck string
	TaskTry   string
}

// ReadJSON is used to Read JSON file from Disk
func (c *ConfigInformation) ReadJSON(path string) error {
	file, _ := os.OpenFile(path, os.O_RDONLY, 0)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(c)
	if err != nil {
		log.Println(err)
	}
	return err
}

// WriteJSON is Used to Write the information as JSON format to Disk
func (c *ConfigInformation) WriteJSON(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(*c)
	if err != nil {
		log.Println(err)
	}
	return err
}
