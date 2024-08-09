package Repositories

import (
	"errors"

	"task/Domain"
)

// TaskRepository is an interface for task repository operations
type TaskRepository interface {
	GetAllTasks() ([]Domain.Task, error)

	GetTaskByID(id int) (*Domain.Task, error)

	CreateTask(task *Domain.Task) error

	UpdateTask(id int, updatedTask *Domain.Task) error

	DeleteTask(id int) error
}

// taskRepository is a concrete implementation of TaskRepository
type taskRepository struct {
	tasks  []Domain.Task
	lastID int
}

// NewTaskRepository creates a new instance of taskRepository
func NewTaskRepository() TaskRepository {
	return &taskRepository{tasks: []Domain.Task{}, lastID: 0}
}

// GetAllTasks retrieves all tasks from the repository
func (r *taskRepository) GetAllTasks() ([]Domain.Task, error) {
	return r.tasks, nil
}

// GetTaskByID retrieves a task by its ID from the repository
func (r *taskRepository) GetTaskByID(id int) (*Domain.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

// CreateTask adds a new task to the repository
func (r *taskRepository) CreateTask(task *Domain.Task) error {
	r.lastID++
	task.ID = r.lastID
	r.tasks = append(r.tasks, *task)
	return nil
}

// UpdateTask updates a task in the repository
func (r *taskRepository) UpdateTask(id int, updatedTask *Domain.Task) error {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks[i] = *updatedTask
			r.tasks[i].ID = id
			return nil
		}
	}
	return errors.New("task not found")
}

// DeleteTask removes a task from the repository
func (r *taskRepository) DeleteTask(id int) error {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
