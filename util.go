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

const PROMPT string = "task-cli>"

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

	fmt.Printf("%s ", PROMPT)
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
		desc = convDesc(argv[1:])
	case "update":
		desc = convDesc(argv[2:])
		fallthrough
	case "mark-in-progress":
		fallthrough
	case "mark-done":
		fallthrough
	case "delete":
		cmd = argv[0]
		tid = convTaskId(argv[1])
	case "quit":
		cmd = argv[0]
	default:
		cmd = "invalid"
	}

	//fmt.Printf("cmd: '%s' %d '%s'", cmd, tid, desc)
	return cmd, tid, desc
}

/*
** Add a task to the list. No need to supply a taskID.
** The function will figure what the next valid unique
** taskID to use.
 */
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

/*
** Delete a task. This function removes the task from
** the list, so it will not be able to be viewed or
** edited after a delete.
 */
func deleteTask(taskid uint) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks = slices.Delete(tasks.Tasks, i, i+1)
			break
		}
	}
}

/*
** Updates a task description on the list indexed by the taskID.
** It also updates the updateAt time of the task.
 */
func updateTask(taskid uint, description string) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks[i].Description = description
			tasks.Tasks[i].UpdatedAt = time.Now().Unix()
			break
		}
	}
}

/*
** Marks a task by changing its status. The only valid
** status's to mark the task is, 'in-progress' or 'done'.
 */
func markTask(taskid uint, status string) {
	for i, v := range tasks.Tasks {
		if v.Id == taskid {
			tasks.Tasks[i].Status = status
			tasks.Tasks[i].UpdatedAt = time.Now().Unix()
			break
		}
	}
}

/*
** List tasks based on the criteria provided. If no critera supplied,
** then all the tasks will be displayed. Valid criteria: "all", "done",
** "in-progress", "not-done", or "todo".
 */
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

/*
** Print all the elements associated with a task.
 */
func printTask(v Task) {
	fmt.Printf("ID: %v\n", v.Id)
	fmt.Printf("    Desc:\t%v\n", v.Description)
	fmt.Printf("    Status:\t%v\n", v.Status)
	fmt.Printf("    CreatedAt:\t%v\n", time.Unix(v.CreatedAt, 0))
	fmt.Printf("    UpdatedAt:\t%v\n\n", time.Unix(v.UpdatedAt, 0))
}

/*
** Convert the string input for taskID to a uint
 */
func convTaskId(strTaskId string) uint {
	u64, err := strconv.ParseUint(strTaskId, 10, 32)
	if err != nil {
		return 0
	}
	return uint(u64)
}

/*
** Convert a slice of strings to a single string.
 */
func convDesc(argv []string) string {
	desc := ""
	spc := ""
	for _, v := range argv {
		desc = desc + spc + v
		spc = " "
	}
	return desc
}
