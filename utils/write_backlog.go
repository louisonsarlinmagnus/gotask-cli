package utils

import (
	"encoding/json"
	"errors"
	"os"

	"gotask-cli/models"
)

const tempPath string = "./tasks.save"

func SaveBacklog(backlog []models.Task, path string) {
	if _, err := os.Stat(tempPath); err == nil {
		err = os.Remove(tempPath)
	} else if errors.Is(err, os.ErrNotExist) {
	} else {
		Check(err)
		os.Exit(1)
	}
	f, err := os.Create(tempPath)
	Check(err)
	defer f.Close()
	backlogJSON, _ := json.MarshalIndent(backlog, "", "\t")
	f.Write(backlogJSON)

	err = os.Rename(tempPath, path)
	Check(err)
}
