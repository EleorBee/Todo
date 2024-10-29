package model

import "time"

type StatusTask string

const (
	TODO        StatusTask = "todo"
	IN_PROGRESS StatusTask = "in-progress"
	DONE        StatusTask = "done"
)

type Task struct {
	//A unique identifier for the task
	Id int
	//A short description of the task
	Description string
	//The status of the task (todo, in-progress, done)
	Status StatusTask
	//The date and time when the task was created
	CreatedAt time.Time
	//The date and time when the task was last updated
	UpdatedAt time.Time
}
