package commands

import (
	"fmt"
	"gotask-cli/models"
	"time"
)

func AddTask(desc string, backlog []models.Task) []models.Task {
	newtask := models.Task{
		Id:          uint(len(backlog)) + 1,
		Description: desc,
		Status:      0,
		Created:     time.Now().Unix(),
		Updated:     -1}

	fmt.Printf("New task \"%s\" added to do @ id=%d", newtask.Description, newtask.Id)
	backlog = append(backlog, newtask)
	return backlog
}
