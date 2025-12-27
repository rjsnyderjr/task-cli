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
		fmt.Printf("Starting new task list\n")
	} else {
		defer jsonFile.Close()
		fmt.Printf("Reading in previous task list\n")
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &tasks)
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
