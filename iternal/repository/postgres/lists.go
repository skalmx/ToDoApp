package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"todoApp/iternal/domain"
)

type ListsRepo struct {
	db *sql.DB
}

func NewListsRepo(db *sql.DB) *ListsRepo {
	return &ListsRepo{
		db: db,
	}
} 

func (l *ListsRepo) Create(ctx context.Context, list domain.List, userId int64) (int64, error) {
	result, err := l.db.Exec("INSERT INTO lists (name) values ($1) RETURNING ID", list.Name)
	if err != nil {
		log.Println("cant create a list")
		return 0, err
	}

	createdId, err := result.LastInsertId()
	if err != nil {
		log.Println("cant get last inserted id")
		return 0, err
	}
	
	_, err = l.db.Exec("INSERT INTO users_lists (user_id, list_id) values ($1, $2)", userId, createdId)
	if err != nil {
		log.Println("cant insert in users_lists table")
		return 0, err
	}
	return createdId, nil
}

func (l *ListsRepo) GetAll(ctx context.Context, userId int64) ([]domain.List, error) {
	rows, err := l.db.Query("SELECT L.id, L.name FROM lists AS L JOIN users_lists AS U ON U.user_id = $1 AND L.id = L.list_id", userId)
	if err != nil {
		log.Println("cant join and select data (getall lists method)")
		return nil, err
	}

	lists := make([]domain.List,0)
	for rows.Next() {
		var list domain.List
		if err := rows.Scan(list.ID, list.Name); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (l *ListsRepo) GetOne(ctx context.Context, userId int64, listId int64) (domain.List, error) {
	var list domain.List
	err := l.db.QueryRow("SELECT L.id, L.name FROM lists AS L JOIN users_lists AS U ON U.user_id = $1 AND L.id = $2 AND L.id = U.list_id", userId, listId).
		Scan(list.ID, list.Name)

	if err == sql.ErrNoRows {
		log.Println("This method doesnt return a row")
		return domain.List{}, err
	}
	if err != nil {
		log.Println("getOne method error while working with lists")
		return domain.List{}, err
	}	

	return list, nil
}

func (l *ListsRepo) Delete(ctx context.Context, userId int64, listId int64) error {
	_, err := l.db.Exec("DELETE FROM lists L USING users_lists U WHERE L.id = U.list_id AND U.user_id=$1 AND L.id=$2", userId, listId)

	if err != nil {
		log.Println("cant delete list")
		return err
	}
	return nil
}

func (l *ListsRepo) Update(ctx context.Context, list domain.ListInput, listId int64 ) error {
	values := make([]string, 0)         // its slice for get all parameters user want to change
	arguments := make([]interface{}, 0) // its slice  with all arguments we need to change
	argNumber := 1                      // counter of arguments

	if list.Name != "" {
		arguments = append(arguments, list.Name)
		values = append(values, fmt.Sprintf("name=$%d", argNumber))
		argNumber++
	}
	arguments = append(arguments, listId)

	queryValues := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id=$%d", queryValues, argNumber)

	_, err := l.db.Exec(query, arguments...)
	if err != nil {
		log.Println("cant update list")
		return err
	}

	return nil
}