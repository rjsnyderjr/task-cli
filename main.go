package main

import (
	"flag"
	"fmt"
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
		case "help":
			fmt.Println("Usage:")
			fmt.Println("  add \"newTaskDescription\"                add a new task to the list")
			fmt.Println("  update taskId \"newTaskDescription\"      update the description of a task")
			fmt.Println("  delete taskId                           remove the task from the list")
			fmt.Println("  mark-in-progress taskId                 change the status of the task to 'in-progress'")
			fmt.Println("  mark-done taskId                        change the status of the task to 'done'")
			fmt.Println("  list [all, done, not-done, in-progress] list the task by their status")
			fmt.Println("  help                                    print this message")
			fmt.Println("  quit                                    save changes and exit program")
		case "quit":
			notDone = false
		default:
			fmt.Println("Invalid command")
		}
	}

	writeOutJson(*fileName)
}
