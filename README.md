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

## USAGE
```
./task-cli <cmd> [arg...]
where,
    cmd - add "TaskDescription"                   add a new task to the list
          update taskId "newTaskDescription"      update the description of a task
          delete taskId                           remove the task from the list
          mark-in-progress taskId                 change the status of the task to 'in-progress'
          mark-done taskId                        change the status of the task to 'done'
          list [all]                              list all tasks
          list todo                               list all todo tasks
          list done                               list all done tasks
          list not-done                           list all not-done tasks
          list in-progress                        list all in-progress tasks
          help                                    print help message
```
