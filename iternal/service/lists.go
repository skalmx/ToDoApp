package service

import (
	"context"
	"errors"
	"todoApp/iternal/domain"
	repository "todoApp/iternal/repository/postgres"
)

type ListService struct {
	repo repository.ListsRepo 
}

func NewListService(repo repository.ListsRepo) *ListService {
	return &ListService {
		repo: repo,
	}
}

func (l *ListService) Create(ctx context.Context, list domain.ListInput, userId int64) (int64, error) {
	if (list == domain.ListInput{}) {
		return 0, errors.New("you cant create a list without name")
	}
	return l.repo.Create(ctx, list, userId)
}

func (l *ListService) GetAll(ctx context.Context, userId int64) ([]domain.List, error) {
	return l.repo.GetAll(ctx, userId)
}

func (l *ListService) GetOne(ctx context.Context, userId int64, listId int64) (domain.List, error) {
	return l.repo.GetOne(ctx, userId, listId)
}

func (l *ListService) Delete(ctx context.Context, userId int64, listId int64) error {
	return l.repo.Delete(ctx, userId, listId)
}

func (l *ListService) Update(ctx context.Context, list domain.ListInput, listId int64) error {
	if (list == domain.ListInput{}) {
		return errors.New("you need to change something to update list")
	}
	return l.repo.Update(ctx, list, listId)
}

