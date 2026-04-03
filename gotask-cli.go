package main

import (
	"fmt"
	"os"
	"strconv"

	"gotask-cli/commands"
	"gotask-cli/utils"
)

const path string = "./tasks.json"

func main() {
	backlog := utils.InitBacklog(path)
	if len(os.Args[1:]) >= 1 {
		switch os.Args[1] {
		case "help":
			commands.Help()
		case "add":
			if len(os.Args[1:]) == 2 {
				backlog = commands.AddTask(os.Args[2], backlog)
			} else {
				fmt.Println("Missing argument")
				commands.Help()
			}
		case "list":
			switch len(os.Args[1:]) {
			case 1:
				commands.ListTasks(backlog)
			case 2:
				switch os.Args[2] {
				case "todo", "doing", "done":
					commands.ListTasksBy(os.Args[2], backlog)
				default:
					fmt.Printf("Unknown status %s", os.Args[2])
				}

			default:
				fmt.Println("Too many arguments !")
			}
		case "progress":
			if len(os.Args[1:]) == 2 {
				id, err := strconv.ParseUint(os.Args[2], 10, 32)
				utils.Check(err)
				if id > 0 && id <= uint64(len(backlog)) {
					commands.ProgressTask(uint(id)-1, backlog)
				} else {
					fmt.Printf("No task found with id=%d, try list to get id", id)
				}
			} else {
				fmt.Println("Missing argument")
				commands.Help()
			}
		case "regress":
			if len(os.Args[1:]) == 2 {
				id, err := strconv.ParseUint(os.Args[2], 10, 32)
				utils.Check(err)
				if id > 0 && id <= uint64(len(backlog)) {
					commands.RegressTask(uint(id)-1, backlog)
				} else {
					fmt.Printf("No task found with id=%d, try list to get id", id)
				}
			} else {
				fmt.Println("Missing argument")
				commands.Help()
			}
		default:
			fmt.Println("NO !")
		}
		utils.SaveBacklog(backlog, path)
	} else {
		fmt.Println("Missing argument")
		commands.Help()

	}
}
