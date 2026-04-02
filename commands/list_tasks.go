package commands

import (
	"gotask-cli/models"
	"gotask-cli/utils"
)

func ListTasks(backlog []models.Task) {
	utils.PrintTasksAsATable(backlog)
}
