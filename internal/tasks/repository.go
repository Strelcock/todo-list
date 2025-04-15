package tasks

import (
	"todoProject/internal/models"
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

func (tr *TaskRepository) Create(task *models.Task) (*models.Task, error) {
	result := tr.Database.DB.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (tr *TaskRepository) GetAll(uid uint) ([]models.Task, error) {
	var tasks []models.Task
	result := tr.Database.DB.Find(&tasks, "user_id = ?", uid)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (tr *TaskRepository) GetDone(uid uint) (*models.Task, error) {
	var task models.Task
	result := tr.Database.DB.Find(&task, "user_id = ? and done = ?", uid, true)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *TaskRepository) GetByName(uid uint, name string) (*models.Task, error) {
	var task models.Task
	result := tr.Database.DB.Find(&task, "user_id = ? and name = ?", uid, name)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (tr *TaskRepository) Mark(task *models.Task, status bool) error {
	result := tr.Database.DB.Model(task).Update("done", status)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TaskRepository) ChangeName(task *models.Task, name string) error {
	result := tr.Database.DB.Model(task).
		Where("name = ?", task.Name).
		Update("name", name)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TaskRepository) Delete(task *models.Task) error {
	result := tr.Database.DB.Delete(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
