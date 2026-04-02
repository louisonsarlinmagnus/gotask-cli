package main

import (
	"fmt"
	"os"

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
			if len(os.Args[1:]) >= 2 {
				backlog = commands.AddTask(os.Args[2], backlog)
			} else {
				fmt.Println("Missing argument")
				commands.Help()
			}
		case "list":
			commands.ListTasks(backlog)
		default:
			fmt.Println("NO !")
		}
		utils.SaveBacklog(backlog, path)
	} else {
		fmt.Println("Missing argument")
		commands.Help()

	}
}
