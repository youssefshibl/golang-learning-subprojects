package src

import "time"

type Task struct {
	id        uint64
	desc      string
	status    TaskStatus
	createdAt time.Time
	updatedAt time.Time
}
