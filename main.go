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

/*
** task-cli [-f taskfile.json]
**	-f taskfile.json - option file to read in at startup. If nor file name is
**	provided in the command line, then the tasks will be saved in tasks.json
**
 */

func main() {
	notDone := true
	changes := false

	fileName := flag.String("f", jsonFileName, "Tasks")
	flag.Parse()

	readInJson(*fileName)

	for notDone {
		cmd, tid, desc := getInput()

		switch cmd {
		case "list":
			listTask(desc)
		case "add":
			addTask(desc)
			changes = true
		case "update":
			updateTask(tid, desc)
			changes = true
		case "delete":
			deleteTask(tid)
			changes = true
		case "mark-done":
			markTask(tid, "done")
			changes = true
		case "mark-in-progress":
			markTask(tid, "in-progress")
			changes = true
		case "help":
			fmt.Println("Usage:")
			fmt.Println("  add \"TaskDescription\"                   add a new task to the list")
			fmt.Println("  update taskId \"newTaskDescription\"      update the description of a task")
			fmt.Println("  delete taskId                           remove the task from the list")
			fmt.Println("  mark-in-progress taskId                 change the status of the task to 'in-progress'")
			fmt.Println("  mark-done taskId                        change the status of the task to 'done'")
			fmt.Println("  list [all, done, not-done,              list the task by their status")
			fmt.Println("        todo, in-progress]")
			fmt.Println("  help                                    print this message")
			fmt.Println("  quit                                    save changes and exit program")
		case "quit":
			notDone = false
		default:
			fmt.Println("Invalid command")
		}

		if changes {
			writeOutJson(*fileName)
			changes = false
		}
	}
}
