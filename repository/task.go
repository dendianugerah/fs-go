package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	task := []entity.Task{}
	err := r.db.WithContext(ctx).Table("tasks").Where("user_id = ?", id).Find(&task).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	err = r.db.WithContext(ctx).Create(&task).Error
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	task := entity.Task{}

	err := r.db.WithContext(ctx).Table("tasks").Where("id = ?", id).Find(&task).Error
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	task := []entity.Task{}
	err := r.db.WithContext(ctx).Table("tasks").Where("category_id = ?", catId).Find(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	err := r.db.WithContext(ctx).Table("tasks").Where("id = ?", task.ID).Updates(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	var task entity.Task
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&task).Error
	if err != nil {
		return err
	}

	return nil
}
