# Task-cli

## DESCRIPTION
A task tracker program for https://roadmap.sh/projects/task-tracker

task-cli is a CLI app to track tasks and manage your todo list. task-cli stores
all tasks in a json file. 

## INSTALLATION
Clone the repository

### To build task-cli:
Windows: C:\> go build -o task-cli.exe main.go json.go util.go
Linux: # go build -o task-cli main.go json.go util.go

### To run task-cli:
./task-cli [-f taskfile.json]

## USAGE
If no -f option is specified, then task-cli will save the default file, 'tasks.json'

task-cli has a help command that will display the command available.

```
task-cli> help
Usage:
  add "TaskDescription"                   add a new task to the list
  update taskId "newTaskDescription"      update the description of a task
  delete taskId                           remove the task from the list
  mark-in-progress taskId                 change the status of the task to 'in-progress'
  mark-done taskId                        change the status of the task to 'done'
  list [all, done, not-done,              list the task by their status
        todo, in-progress]
  help                                    print this message
  quit                                    save changes and exit program
 ```  
