package tasks

import "todoProject/pkg/db"

type TaskRepository struct {
	Database *db.Db
}

func NewTaskRepository(DB *db.Db) *TaskRepository {
	return &TaskRepository{
		Database: DB,
	}
}

func (tr *TaskRepository) Create(task *Task) (*Task, error) {
	result := tr.Database.DB.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (tr *TaskRepository) GetByName(name string) (*Task, error) {
	var task Task
	result := tr.Database.DB.First(&task, "name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}
