package domain

import "time"

type Task struct {
	ID           uint64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ExpiresAt    time.Time `json:"expiresAt"`
	UserID		 int64	   `json:"userId"`
}

type TaskInput struct {
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ExpiresAt    time.Time `json:"expiresAt"`
	UserID		 uint64	   `json:"userId"`
}