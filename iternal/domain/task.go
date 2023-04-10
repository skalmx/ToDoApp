package domain

import "time"

type Task struct {
	ID           uint64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ExpiresAt    time.Time `json:"expiresAt"`
	ListID		 uint64	   `json:"listId"`
}

type TaskInput struct {
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ExpiresAt    time.Time `json:"expiresAt"`
	ListID		 uint64	   `json:"listId"`
}