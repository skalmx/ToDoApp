package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"todoApp/iternal/domain"

	"golang.org/x/net/context"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (u *UsersRepo) CreateUser(ctx context.Context, user domain.User) (int64, error) {
	result, err := u.db.Exec("INSERT INTO users (name, email, phone, password, registeredAt) values ($1, $2, $3, $4, $5) RETURNING ID",
		user.Name, user.Email, user.Phone, user.Phone, user.RegisteredAt)

	if err != nil {
		log.Println("cant create a user in db")
		return 0, err
	}

	id, err := result.LastInsertId()

	return id, err
}

func (u *UsersRepo) UpdateUser(ctx context.Context, user domain.UserInput, id int64) error {
	values := make([]string, 0)         // its slice for get all parameters user want to change
	arguments := make([]interface{}, 0) // its slice  with all arguments we need to change
	argNumber := 1                      // counter of arguments

	if user.Name != "" {
		arguments = append(arguments, user.Name)
		values = append(values, fmt.Sprintf("name=$%d", argNumber))
		argNumber++
	}

	if user.Email != "" {
		arguments = append(arguments, user.Email)
		values = append(values, fmt.Sprintf("email=$%d", argNumber))
		argNumber++
	}

	if user.Phone != "" {
		arguments = append(arguments, user.Phone)
		values = append(values, fmt.Sprintf("phone=$%d", argNumber))
		argNumber++
	}

	if user.Password != "" {
		arguments = append(arguments, user.Password)
		values = append(values, fmt.Sprintf("password=$%d", argNumber))
		argNumber++
	}

	arguments = append(arguments, id)
	queryValues := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", queryValues, argNumber) // get a query like UPDATE {} SET {smth}="123" .... WHERE ID = id

	_, err := u.db.Exec(query, arguments...)
	return err
}
