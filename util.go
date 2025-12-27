package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const prompt string = "task-cli>"

func getInput() (string, uint, string) {
	var argv []string = make([]string, 0)
	var cmd string = ""
	var tid uint = 0
	var desc string = ""

	/*
	** Valid commands:
	**	add "newTaskDescription"
	**	update taskId "newTaskDescription"
	**	delete taskId
	**	mark-in-progress taskId
	**	mark-done taskId
	**	list [all, done, not-done, in-progress]
	**	help
	**	quit
	 */

	fmt.Printf("%s ", prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fields := strings.Fields(strings.TrimSpace(input))

	for _, field := range fields {
		argv = append(argv, strings.Trim(field, "\""))
	}

	switch argv[0] {
	case "help":
		cmd = argv[0]
	case "add":
		fallthrough
	case "list":
		cmd = argv[0]
		desc = ""
		spc := ""
		for i, v := range argv {
			if i > 0 {
				desc = desc + spc + v
				spc = " "
			}
		}
	case "update":
		cmd = argv[0]
		u64, err := strconv.ParseUint(argv[1], 10, 32)
		if err != nil {
			break
		}
		tid = uint(u64)
		desc = ""
		spc := ""
		for i, v := range argv {
			if i > 1 {
				desc = desc + spc + v
				spc = " "
			}
		}
	case "mark-in-progress":
		fallthrough
	case "mark-done":
		fallthrough
	case "delete":
		cmd = argv[0]
		u64, err := strconv.ParseUint(argv[1], 10, 32)
		if err != nil {
			break
		}
		tid = uint(u64)
	case "quit":
		cmd = argv[0]
	default:
		cmd = "invalid"
	}

	return cmd, tid, desc
}

func addTask(description string) {
	if taskId == 0 {
		var lastId uint = 0
		for _, v := range tasks.Tasks {
			lastId = v.Id
		}
		taskId = lastId
	}
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
			tasks.Tasks[i].UpdatedAt = time.Now().Unix()
			break
		}
	}
}

func markTask(taskid uint, status string) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks[i].Status = status
			tasks.Tasks[i].UpdatedAt = time.Now().Unix()
			break
		}
	}
}

func listTask(action string) {
	if action == "" {
		action = "all"
	}
	if action != "all" && action != "done" && action != "in-progress" && action != "todo" && action != "not-done" {
		fmt.Println("Invalid list option")
		return
	}
	fmt.Println("=================================")
	for _, v := range tasks.Tasks {
		if action == "not-done" {
			if v.Status != "done" {
				printTask(v)
			}
		} else {
			if action == "all" || v.Status == action {
				printTask(v)

			}
		}
	}
}

func printTask(v Task) {
	fmt.Printf("ID: %v\n", v.Id)
	fmt.Printf("    Desc:\t%v\n", v.Description)
	fmt.Printf("    Status:\t%v\n", v.Status)
	fmt.Printf("    CreatedAt:\t%v\n", time.Unix(v.CreatedAt, 0))
	fmt.Printf("    UpdatedAt:\t%v\n\n", time.Unix(v.UpdatedAt, 0))
}

func getNewTaskId() uint {

	return taskId
}
