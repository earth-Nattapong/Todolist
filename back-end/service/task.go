package service

import "time"

type TaskResponse struct {
	ID         int       `json:"id"`
	Taskname   string    `json:"taskname"`
	Title      string    `json:"title"`
	Time       time.Time `json:"time"`
	Check      string    `json:"check"`
	Importance int       `json:"importance"`
}
type TaskRequest struct {
	Taskname   string `json:"taskname"`
	Title      string `json:"title"`
	Check      string `json:"check"`
	Importance int    `json:"importance"`
}

// เลือกแค่ต้องการให้แสดงผล
type TaskService interface {
	GetTaskAllCos() ([]TaskResponse, error)
	GetTaskByIdCos(id int) (*TaskResponse, error)
	CreateTaskCos(task TaskRequest) (*TaskResponse, error)
	UpdateTaskCos(id int, task TaskRequest) (*TaskResponse, error)
	DeleteTaskCos(id int) error
}
