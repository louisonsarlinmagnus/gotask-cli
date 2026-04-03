package utils

import (
	"encoding/json"
	"errors"
	"gotask-cli/models"
	"os"
)

func InitBacklog(path string) []models.Task {
	if _, err := os.Stat(path); err == nil {

	} else if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(path)
		Check(err)
		defer f.Close()
		f.Write([]byte("[]"))

	} else {
		Check(err)
		os.Exit(1)
	}
	return loadBacklog(path)
}

func loadBacklog(path string) []models.Task {
	var backlog []models.Task
	content, err := os.ReadFile(path)
	Check(err)
	err = json.Unmarshal([]byte(content), &backlog)
	Check(err)
	return backlog
}
