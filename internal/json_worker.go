package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func AddJsonData(fileName string, data string) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Creating " + fileName)

			file = []byte("[]")
		} else {
			log.Fatal(err)
		}
	}

	var jsonData []string

	if err := json.Unmarshal(file, &jsonData); err != nil {
		log.Fatal(err)
	}

	updatedData, err := json.MarshalIndent(append(jsonData, data), "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(fileName, updatedData, 0644); err != nil {
		log.Fatal(err)
	}
}
