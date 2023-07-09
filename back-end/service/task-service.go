package service

import (
	"log"
	"time"

	"example.com/hello/repository"
)

type taskService struct {
	taskRepo repository.TaskRespository
}

func NewTaskService(taskRepo repository.TaskRespository) taskService {
	return taskService{taskRepo: taskRepo}
}

func (s taskService) GetTaskAllCos() ([]TaskResponse, error) {
	tasks, err := s.taskRepo.GetTaskAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	taskResponses := []TaskResponse{}
	for _, task := range tasks {
		taskResponse := TaskResponse{
			ID:         task.ID,
			Taskname:   task.Taskname,
			Title:      task.Title,
			Time:       task.Time,
			Check:      task.Check,
			Importance: task.Importance,
		}
		taskResponses = append(taskResponses, taskResponse)
	}
	return taskResponses, nil
}

func (s taskService) GetTaskByIdCos(id int) (*TaskResponse, error) {
	tasks, err := s.taskRepo.GetTaskById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	taskResponses := TaskResponse{
		ID:         tasks.ID,
		Taskname:   tasks.Taskname,
		Title:      tasks.Title,
		Time:       tasks.Time,
		Check:      tasks.Check,
		Importance: tasks.Importance,
	}

	return &taskResponses, nil
}

func (s taskService) CreateTaskCos(task TaskRequest) (*TaskResponse, error) {
	Newtask := repository.Task{
		Taskname:   task.Taskname,
		Title:      task.Title,
		Time:       time.Now(),
		Check:      task.Check,
		Importance: task.Importance,
	}
	tasks, err := s.taskRepo.CreateTask(Newtask)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	taskResponses := TaskResponse{
		ID:         tasks.ID,
		Taskname:   tasks.Taskname,
		Title:      tasks.Title,
		Time:       tasks.Time,
		Check:      tasks.Check,
		Importance: tasks.Importance,
	}

	return &taskResponses, nil

}

func (s taskService) UpdateTaskCos(id int, task TaskRequest) (*TaskResponse, error) {
	Newtask := repository.Task{
		ID:         id,
		Taskname:   task.Taskname,
		Title:      task.Title,
		Time:       time.Now(),
		Check:      task.Check,
		Importance: task.Importance,
	}

	tasks, err := s.taskRepo.UpdateTask(id, Newtask)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	taskResponses := TaskResponse{
		ID:         tasks.ID,
		Taskname:   tasks.Taskname,
		Title:      tasks.Title,
		Time:       tasks.Time,
		Check:      tasks.Check,
		Importance: tasks.Importance,
	}

	return &taskResponses, nil
}

func (s taskService) DeleteTaskCos(id int) error {
	err := s.taskRepo.DeleteTask(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
