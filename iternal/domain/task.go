package domain

import "time"

type Task struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

type TaskInput struct {
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ExpiresAt    time.Time `json:"expiresAt"`
}