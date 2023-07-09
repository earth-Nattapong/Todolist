package repository

import "time"

type Task struct {
	ID         int
	Taskname   string
	Title      string
	Time       time.Time
	Check      string
	Importance int
}

type TaskRespository interface {
	GetTaskAll() ([]Task, error)
	GetTaskById(id int) (*Task, error)
	CreateTask(task Task) (*Task, error)
	UpdateTask(id int, task Task) (*Task, error)
	DeleteTask(id int) error
}
