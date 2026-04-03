package commands

import (
	"fmt"
	"gotask-cli/models"
	"gotask-cli/utils"
)

func ListTasksBy(status string, backlog []models.Task) {
	var backlogToDisplay []models.Task
	var statusID models.Status
	for i, v := range models.StatusName {
		if v == status {
			statusID = i
		}
	}
	for _, v := range backlog {
		if v.Status == models.Status(statusID) {
			backlogToDisplay = append(backlogToDisplay, v)
		}
	}
	if len(backlogToDisplay) > 0 {
		utils.PrintTasksAsATable(backlogToDisplay)
	} else {
		fmt.Printf("No tasks with the status [%s]", status)
	}

}
