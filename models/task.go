package models

type Status int

const (
	todo Status = iota
	doing
	done
)

var StatusName = map[Status]string{
	todo:  "todo",
	doing: "doing",
	done:  "done",
}

type Task struct {
	Id          uint
	Description string
	Status      Status
	Created     int64
	Updated     int64
}
