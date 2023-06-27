
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	User            string `json:"user"`
	Password        string `json:"password"`
	DBName          string `json:"dbName"`
	PluginType      string `json:"pluginType"`
	SourceDirectory string `json:"sourceDirectory"`
}

func main() {
	jsonFilePath := "/home/swati/Documents/shared_folder/cred.json"

	// Read the JSON file contents
	jsonBytes, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse the JSON contents into the Config struct
	var config Config
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Validate and process the source and destination directory paths
	if config.SourceDirectory == "" {
		fmt.Println("Source directory path is missing in the JSON file.")
		return
	}

	// Get the source directory path from the config
	sourceDirPath := config.SourceDirectory

	// Use the sourceDirPath as needed
	fmt.Println("Source Directory:", sourceDirPath)
}
