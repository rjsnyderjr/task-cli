package main

import (
	"fmt"
	"os"
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
** task-cli
**   add "TaskDescription"                   add a new task to the list
**   update taskId "newTaskDescription"      update the description of a task
**   delete taskId                           remove the task from the list
**   mark-in-progress taskId                 change the status of the task to 'in-progress'
**   mark-done taskId                        change the status of the task to 'done'
**   list [all]                              list all tasks
**   list done                               list tasks where status is 'done'
**   list not-done                           list tasks where status is 'not done'
**   list todo                               list tasks where status is 'todo'
**   list in-progress                        list tasks where status is 'in-progress'
**   help                                    show list of commands
**
 */

func main() {
	argc := len(os.Args)

	if argc < 2 || argc > 4 {
		fmt.Println("Invalid number of parameters")
		os.Exit(1)
	}

	readInJson(jsonFileName)

	progName := os.Args[0]

	cmd := os.Args[1]

	switch cmd {
	case "list":
		switch argc {
		case 2:
			listTask("")
		case 3:
			listTask(os.Args[2])
		default:
			fmt.Println("Invalid number of parameters")
			os.Exit(1)
		}
	case "add":
		if argc != 3 {
			fmt.Println("Invalid number of parameters")
			os.Exit(1)
		}
		addTask(os.Args[2])
	case "update":
		if argc != 4 {
			fmt.Println("Invalid number of parameters")
			os.Exit(1)
		}
		updateTask(os.Args[2], os.Args[3])
	case "delete":
		if argc != 3 {
			fmt.Println("Invalid number of parameters")
			os.Exit(1)
		}
		deleteTask(os.Args[2])
	case "mark-done":
		if argc != 3 {
			fmt.Println("Invalid number of parameters")
			os.Exit(1)
		}
		markTask(os.Args[2], "done")
	case "mark-in-progress":
		if argc != 3 {
			fmt.Println("Invalid number of parameters")
			os.Exit(1)
		}
		markTask(os.Args[2], "in-progress")
	case "help":
		fmt.Println("Usage:")
		fmt.Printf("%s  add \"TaskDescription\"                   add a new task to the list\n", progName)
		fmt.Printf("%s  update taskId \"newTaskDescription\"      update the description of a task\n", progName)
		fmt.Printf("%s  delete taskId                           remove the task from the list\n", progName)
		fmt.Printf("%s  mark-in-progress taskId                 change the status of the task to 'in-progress'\n", progName)
		fmt.Printf("%s  mark-done taskId                        change the status of the task to 'done'\n", progName)
		fmt.Printf("%s  list [all]                              list all tasks\n", progName)
		fmt.Printf("%s  list todo                               list all todo tasks\n", progName)
		fmt.Printf("%s  list done                               list all done tasks\n", progName)
		fmt.Printf("%s  list not-done                           list all not-done tasks\n", progName)
		fmt.Printf("%s  list in-progress                        list all in-progress tasks\n", progName)
		fmt.Printf("%s  help                                    print this message", progName)
	default:
		fmt.Println("Invalid command")
	}

	writeOutJson(jsonFileName)
	os.Exit(0)
}
