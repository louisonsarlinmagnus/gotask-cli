package commands

import (
	"fmt"
	"gotask-cli/models"
	"time"
)

func ProgressTask(id uint, backlog []models.Task) []models.Task {
	if backlog[id].Status >= 0 && backlog[id].Status < 2 {
		backlog[id].Status++
		backlog[id].Updated = time.Now().Unix()
		fmt.Printf("[%s] task \"%s\" progress to [%s]", models.StatusName[backlog[id].Status-1], backlog[id].Description, models.StatusName[backlog[id].Status])
	} else {
		fmt.Printf("Can't progress task %s", backlog[id].Description)
	}
	return backlog
}
