package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const jsonFileName string = "./tasks.json"
const prompt string = "task-cli "

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Id          uint   `json:"id"`          // unique ID 1-
	Description string `json:"description"` // task description
	Status      string `json:"status"`      // status: 'todo', 'in-progress', 'done'
	CreatedAt   int64  `json:"createdat"`   // timestamp when task was created
	UpdatedAt   int64  `json:"updatedat"`   // timestamp when task was last updated
}

var taskId uint = 0
var tasks Tasks
var argc int = 0
var args []string

func main() {
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Printf("Starting new task list\n")
	} else {
		defer jsonFile.Close()
		fmt.Printf("Reading in previous task list\n")
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &tasks)
	}

	notDone := true

	for notDone {
		argc = getInput()

		switch args[0] {
		case "list":
			switch argc {
			case 1:
				listTask("all")
			case 2:
				listTask(args[1])
			}
		case "add":
			if argc < 2 {
				break
			}
			desc := ""
			spc := ""
			for i, v := range args {
				if i > 0 {
					desc = desc + spc + v
					spc = " "
				}
			}
			addTask(strings.Trim(desc, "\""))
		case "update":
			if argc < 3 {
				break
			}
			u64, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				break
			}
			desc := ""
			spc := ""
			for i, v := range args {
				if i > 1 {
					desc = desc + spc + v
					spc = " "
				}
			}
			updateTask(uint(u64), strings.Trim(desc, "\""))
		case "delete":
			if argc < 2 {
				break
			}
			u64, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				break
			}
			deleteTask(uint(u64))
		case "mark-done":
			if argc < 2 {
				break
			}
			u64, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				break
			}
			markTask(uint(u64), "done")
		case "mark-in-progress":
			if argc < 2 {
				break
			}
			u64, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				break
			}
			markTask(uint(u64), "in-progress")
		case "quit":
			notDone = false
		}
	}

	jsonData, err := json.Marshal(&tasks)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(jsonFileName, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getInput() int {
	args = args[:0]
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fields := strings.Fields(strings.TrimSpace(input))

	for _, field := range fields {
		args = append(args, field)
	}
	return len(args)
}

func addTask(description string) {
	taskId++
	now := time.Now().Unix()
	var task = Task{
		Id:          taskId,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks.Tasks = append(tasks.Tasks, task)
	fmt.Printf("Task added successfully (ID: %v)\n", taskId)
}

func deleteTask(taskid uint) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks = slices.Delete(tasks.Tasks, i, i+1)
			break
		}
	}
}

func updateTask(taskid uint, description string) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks[i].Description = description
			break
		}
	}
}

func markTask(taskid uint, status string) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks[i].Status = status
			break
		}
	}
}

func listTask(action string) {
	if action != "all" && action != "done" && action != "in-progress" && action != "todo" {
		return
	}
	fmt.Println("=================================")
	for _, v := range tasks.Tasks {
		if action == "all" || v.Status == action {
			fmt.Printf("ID: %v\n", v.Id)
			fmt.Printf("    Desc:\t%v\n", v.Description)
			fmt.Printf("    Status:\t%v\n", v.Status)
			fmt.Printf("    CreatedAt:\t%v\n", time.Unix(v.CreatedAt, 0))
			fmt.Printf("    UpdatedAt:\t%v\n\n", time.Unix(v.UpdatedAt, 0))
		}
	}
}
