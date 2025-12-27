package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func readInJson(filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		taskId = 0
	} else {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		err = json.Unmarshal(byteValue, &tasks)
		if err != nil {
			fmt.Printf("Error reading json file %s\n", filename)
			log.Fatal(err)
		}
	}
	//fmt.Printf("%d tasks read in\n", len(tasks.Tasks))
}

func writeOutJson(filename string) {
	//fmt.Printf("%d tasks saved\n", len(tasks.Tasks))
	jsonData, err := json.Marshal(&tasks)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
