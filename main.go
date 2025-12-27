package main

import (
	"flag"
)

const jsonFileName string = "tasks.json"

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

func main() {
	fileName := flag.String("f", jsonFileName, "Tasks")
	flag.Parse()

	readInJson(*fileName)

	notDone := true

	for notDone {
		cmd, tid, desc := getInput()

		switch cmd {
		case "list":
			listTask(desc)
		case "add":
			addTask(desc)
		case "update":
			updateTask(tid, desc)
		case "delete":
			deleteTask(tid)
		case "mark-done":
			markTask(tid, "done")
		case "mark-in-progress":
			markTask(tid, "in-progress")
		case "quit":
			notDone = false
		}
	}

	writeOutJson(*fileName)
}
