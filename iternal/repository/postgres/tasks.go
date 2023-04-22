package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"todoApp/iternal/domain"
)

type TasksRepo struct {
	db *sql.DB
}

func NewTasksRepo(db *sql.DB) *TasksRepo {
	return &TasksRepo{db}
}

func (t *TasksRepo) Create(ctx context.Context, task domain.TaskInput, listId int64) (int64, error) {
	result, err := t.db.Exec("INSERT INTO tasks (name, status, expiresAt) values ($1, $2, $3) RETURNING ID",
					task.Name, task.Status, task.ExpiresAt)
	
	if err != nil {
		log.Println("cant create a task")
		return 0, err
	}
	createdId, err := result.LastInsertId()
	if err != nil {
		log.Println("cant get last inserted id")
		return 0, err
	}
	_, err = t.db.Exec("INSERT INTO lists_tasks (list_id, task_id) values ($1, $2)",
						listId, createdId)
	
	if err != nil {
		log.Println("cant insert in lists_tasks table")
		return 0, err
	}

	return createdId, nil //todo add transaction
}

func (t *TasksRepo) GetAll(ctx context.Context, listId int64) ([]domain.Task, error) {
	rows, err := t.db.Query("SELECT T.id, T.name, T.status, T.expiresAt FROM tasks AS T JOIN lists_tasks AS L ON L.list_id = $1 AND T.id = L.task_id", listId)
	if err != nil {
		log.Println("cant join and select data (getall tasks method)")
		return nil, err
	}

	tasks := make([]domain.Task,0)
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(task.ID, task.Name, task.Status, task.ExpiresAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *TasksRepo) GetOne(ctx context.Context, listId int64, taskId int64) (domain.Task, error) {
	var task domain.Task
	err := t.db.QueryRow("SELECT T.id, T.name, T.status, T.expiresAt FROM tasks AS T JOIN lists_tasks AS L ON L.list_id = $1 AND T.id = $2 AND T.id = L.task_id", listId, taskId).
		Scan(task.ID, task.Name, task.Status, task.ExpiresAt)

	if err == sql.ErrNoRows {
		log.Println("This method doesnt return a row")
		return domain.Task{}, err
	}
	if err != nil {
		log.Println("getOne method error while working with tasks")
		return domain.Task{}, err
	}	

	return task, nil
}

func (t *TasksRepo) Delete(ctx context.Context, listId int64, taskId int64) error {
	_, err := t.db.Exec("DELETE FROM tasks t USING lists_tasks l WHERE t.id=l.task_id AND l.list_id=$1 AND t.id=$2", listId, taskId)

	if err != nil {
		log.Println("cant delete task")
		return err
	}
	return nil
}

func (t *TasksRepo) Update(ctx context.Context, task domain.TaskInput, taskId int64) error {
	values := make([]string, 0) // its slice for get all parameters user want to change
	arguments := make([]interface{}, 0) // its slice  with all arguments we need to change
	argNumber := 1 // counter of arguments 

	if task.Name != "" {
		arguments = append(arguments, task.Name)
		values = append(values, fmt.Sprintf("name=$%d", argNumber))
		argNumber++
	}
	if task.Status != "" {
		arguments = append(arguments, task.Status)
		values = append(values, fmt.Sprintf("status=$%d", argNumber))
		argNumber++
	}
	if !task.ExpiresAt.IsZero() {
		arguments = append(arguments, task.ExpiresAt)
		values = append(values, fmt.Sprintf("expiresAt=$%d", argNumber))
		argNumber++
	}
	
	arguments = append(arguments, taskId)

	queryValues := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id=$%d", queryValues, argNumber)

	_, err := t.db.Exec(query, arguments...)
	if err != nil {
		log.Println("cant update task in tasks table")
		return err
	}

	return nil
}