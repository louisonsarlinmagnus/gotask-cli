package commands

import (
	"fmt"
	"gotask-cli/models"
	"time"
)

func RegressTask(id uint, backlog []models.Task) []models.Task {
	if backlog[id].Status >= 1 && backlog[id].Status <= 2 {
		backlog[id].Status--
		backlog[id].Updated = time.Now().Unix()
		fmt.Printf("[%s] task \"%s\" regress to [%s]", models.StatusName[backlog[id].Status+1], backlog[id].Description, models.StatusName[backlog[id].Status])
	} else {
		fmt.Printf("Can't regress task %s", backlog[id].Description)
	}
	return backlog
}
