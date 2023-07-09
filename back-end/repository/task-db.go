package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type taskRepositoryDB struct {
	db *gorm.DB
}

func NewTaskRepositoryDB(db *gorm.DB) taskRepositoryDB {
	return taskRepositoryDB{db: db}
}

func (r taskRepositoryDB) GetTaskAll() ([]Task, error) {
	tasks := []Task{}
	query := r.db.Find(&tasks)
	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, query.Error
	}
	return tasks, nil
}

func (r taskRepositoryDB) GetTaskById(id int) (*Task, error) {
	tasks := Task{}
	query := r.db.Find(&tasks, id)
	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, query.Error
	}
	return &tasks, nil
}

func (r taskRepositoryDB) CreateTask(task Task) (*Task, error) {
	query := r.db.Create(&task)
	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, query.Error
	}
	return &task, nil
}

func (r taskRepositoryDB) UpdateTask(id int, task Task) (*Task, error) {
	query := r.db.Model(&Task{}).Where("id=?", id).Updates(&task)
	if query.Error != nil {
		fmt.Println(query.Error)
		return nil, query.Error
	}
	return r.GetTaskById(id)
}

func (r taskRepositoryDB) DeleteTask(id int) error {
	query := r.db.Delete(&Task{}, id)
	if query.Error != nil {
		fmt.Println(query.Error)
		return query.Error
	}
	return nil
}
