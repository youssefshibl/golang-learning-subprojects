package src

type TaskStatus string

const (
	Done    TaskStatus = "Done"
	Pending TaskStatus = "Pending"
)

type CommandActions string

const (
	Add CommandActions = "add"
)

type SomeConfig string

const (
	TaskFileName SomeConfig = "task.json"
)
