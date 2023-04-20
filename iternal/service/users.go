package service

import (
	"context"
	"todoApp/iternal/domain"
	repository "todoApp/iternal/repository/postgres"
)

type UserService struct {
	repo repository.UsersRepo
}

func NewUserService(repo repository.UsersRepo) *UserService {
	return &UserService{
		repo : repo,
	}
}

func (u *UserService) Create(ctx context.Context, user domain.User ) (int64, error) {
	return u.repo.Create(ctx, user)
}

func (u *UserService) Update(ctx context.Context, user domain.UserInput, id int64) error {
	if (user == domain.UserInput{}) {
		return nil	
	}
	return u.repo.Update(ctx, user, id)
}