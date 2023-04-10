package domain

import "time"

type User struct {
	ID           uint64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registeredAt"`
}