package domain

import "time"

type Task struct {
	ID string
	Title string
	Completed bool
	CreatedAt time.Time
}