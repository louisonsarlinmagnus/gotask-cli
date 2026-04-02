package models

type status int

const (
	todo status = iota
	doing
	done
)

var StatusName = map[status]string{
	todo:  "todo",
	doing: "doing",
	done:  "done",
}

type Task struct {
	Id          uint
	Description string
	Status      status
	Created     int64
	Updated     int64
}
