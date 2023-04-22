package service

import (
	"context"
	"errors"
	"todoApp/iternal/domain"
	repository "todoApp/iternal/repository/postgres"
)

type TaskService struct {
	repo repository.TasksRepo
}

func NewTaskService(repo repository.TasksRepo) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) Create(ctx context.Context, task domain.TaskInput, listId int64) (int64, error) {
	if task.Name == "" {
		return 0, errors.New("cant create a task without name")
	}

	if task.Status == "" {
		task.Status = "New"
	}

	if task.ExpiresAt.IsZero() {
		return 0, errors.New("cant create a task without deadline") // todo u can create a task without deadline
	}

	return t.repo.Create(ctx, task, listId)
}

func (t *TaskService) GetAll(ctx context.Context, listId int64) ([]domain.Task, error) {
	return t.repo.GetAll(ctx, listId)
}

func (t *TaskService) GetOne(ctx context.Context, listId int64, taskId int64) (domain.Task, error) {
	return t.repo.GetOne(ctx, listId, taskId)
}

func (t *TaskService) Delete(ctx context.Context, listId int64, taskId int64) error {
	return t.repo.Delete(ctx, listId, taskId)
}

func (t *TaskService) Update(ctx context.Context, task domain.TaskInput, taskId int64) error {
	if (task == domain.TaskInput{}) {
		return errors.New("you need to change something to update task")
	}
	return t.repo.Update(ctx, task, taskId)
}