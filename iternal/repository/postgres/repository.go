package repository

import (
	"context"
	"database/sql"
	"todoApp/iternal/domain"
)

type Users interface {
	Create(ctx context.Context, user domain.User) (int64, error)
	Update(ctx context.Context, user domain.UserInput, id int64) error
}

type Lists interface {

}

type Tasks interface {
	Create(ctx context.Context, task domain.TaskInput, listId int64) (int64, error)
	GetAll(ctx context.Context, listId int64) ([]domain.Task, error)
	GetOne(ctx context.Context, listId int64, taskId int64) (domain.Task, error)
	Delete(ctx context.Context, listId int64, taskId int64) error
	Update(ctx context.Context, input domain.TaskInput, taskId int64) error
}
type Repositories struct {
	Tasks Tasks
	Lists Lists
	Users Users
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Tasks: NewTasksRepo(db),
		Lists: NewListsRepo(db),
		Users: NewUsersRepo(db),
	}
}