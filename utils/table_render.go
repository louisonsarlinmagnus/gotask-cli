package utils

import (
	"gotask-cli/models"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintTasksAsATable(tasks []models.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Description", "Status", "Created", "Updated"})
	for _, v := range tasks {
		var updated string
		if v.Updated == -1 {
			updated = ""
		} else {
			updated = time.Unix(v.Updated, 0).Format(time.DateTime)
		}
		t.AppendRow([]interface{}{
			v.Id,
			v.Description,
			models.StatusName[v.Status],
			time.Unix(v.Created, 0).Format(time.DateTime),
			updated,
		})
	}
	t.Render()
}
