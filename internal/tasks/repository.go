package tasks

import (
	"todoProject/pkg/db"
)

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

func (tr *TaskRepository) GetAll(uid uint) ([]Task, error) {
	var tasks []Task
	result := tr.Database.DB.Find(&tasks, "user_id = ?", uid)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (tr *TaskRepository) GetDone(uid uint) (*Task, error) {
	var task Task
	result := tr.Database.DB.Find(&task, "user_id = ? and done = ?", uid, true)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *TaskRepository) GetByName(uid uint, name string) (*Task, error) {
	var task Task
	result := tr.Database.DB.Find(&task, "user_id = ? and name = ?", uid, name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *TaskRepository) Mark(task *Task, status bool) error {
	result := tr.Database.DB.Model(task).Update("done", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TaskRepository) ChangeName(task *Task, name string) error {
	result := tr.Database.DB.Model(task).
		Where("name = ?", task.Name).
		Update("name", name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TaskRepository) Delete(task *Task) error {
	result := tr.Database.DB.Delete(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
