package Usecases

import (
	"task/Domain"
	"task/Repositories"
)

// a use case for handling tasks
type ITaskUseCase interface {
	GetAllTasks() ([]Domain.Task, error)

	CreateTask(task *Domain.Task) error

	GetTaskByID(id int) (*Domain.Task, error)

	UpdateTask(id int, task *Domain.Task) error

	DeleteTask(id int) error
}

// TaskUseCase is a use case for handling tasks
type TaskUseCase struct {
	TaskRepo Repositories.TaskRepository
}

// GetAllTasks gets all tasks
func (uc *TaskUseCase) GetAllTasks() ([]Domain.Task, error) {
	return uc.TaskRepo.GetAllTasks()
}

// GetTaskByID gets a task by ID
func (uc *TaskUseCase) GetTaskByID(id int) (*Domain.Task, error) {
	return uc.TaskRepo.GetTaskByID(id)
}

// CreateTask creates a new task
func (uc *TaskUseCase) CreateTask(task *Domain.Task) error {
	return uc.TaskRepo.CreateTask(task)
}

// UpdateTask updates a task by ID
func (uc *TaskUseCase) UpdateTask(id int, updatedTask *Domain.Task) error {
	return uc.TaskRepo.UpdateTask(id, updatedTask)
}

// DeleteTask deletes a task by ID
func (uc *TaskUseCase) DeleteTask(id int) error {
	return uc.TaskRepo.DeleteTask(id)
}
