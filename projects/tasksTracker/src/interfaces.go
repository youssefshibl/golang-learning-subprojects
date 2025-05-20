package src

import "time"

type Task struct {
	Id        int64      `json:"id"`
	Desc      string     `json:"description"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
