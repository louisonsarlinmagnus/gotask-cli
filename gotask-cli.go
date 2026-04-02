package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Status int

const (
	todo Status = iota
	doing
	done
)

const path string = "./tasks.json"
const tempPath string = "./tasks.save"

var statusName = map[Status]string{
	todo:  "todo",
	doing: "doing",
	done:  "done",
}

type task struct {
	Id          uint
	Description string
	Status      Status
	Created     int64
	Updated     int64
}

var backlog []task

func initBacklog() {
	if _, err := os.Stat(path); err == nil {
		loadBacklog()

	} else if errors.Is(err, os.ErrNotExist) {
		// path does not exist
		f, err := os.Create(path)
		check(err)
		defer f.Close()
		loadBacklog()

	} else {
		check(err)
		os.Exit(1)
	}
}

func loadBacklog() {
	content, err := os.ReadFile(path)
	check(err)
	err = json.Unmarshal([]byte(content), &backlog)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printTasksAsATable(tasks []task) {
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
			statusName[v.Status],
			time.Unix(v.Created, 0).Format(time.DateTime),
			updated,
		})
	}
	t.Render()
}

func addTask(desc string) {
	newtask := task{
		Id:          uint(len(backlog)) + 1,
		Description: desc,
		Status:      0,
		Created:     time.Now().Unix(),
		Updated:     -1}

	fmt.Printf("New task \"%s\" added to do @ id=%d", newtask.Description, newtask.Id)
	backlog = append(backlog, newtask)
}

func saveBacklog() {
	if _, err := os.Stat(tempPath); err == nil {
		err = os.Remove(tempPath)
	} else if errors.Is(err, os.ErrNotExist) {
	} else {
		check(err)
		os.Exit(1)
	}
	f, err := os.Create(tempPath)
	check(err)
	defer f.Close()
	backlogJSON, _ := json.MarshalIndent(backlog, "", "\t")
	f.Write(backlogJSON)

	err = os.Rename(tempPath, path)
	check(err)
}

func main() {
	initBacklog()
	if len(os.Args[1:]) >= 1 {
		switch os.Args[1] {
		case "help":
			fmt.Println("S.O.S!")
		case "add":
			if len(os.Args[1:]) >= 2 {
				addTask(os.Args[2])
			} else {
				fmt.Println("Missing argument")
			}
		case "list":
			printTasksAsATable(backlog)
		default:
			fmt.Println("NO !")
		}
		saveBacklog()
	} else {
		fmt.Println("Missing argument")
	}
}
