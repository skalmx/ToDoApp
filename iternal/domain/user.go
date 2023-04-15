package domain

import "time"

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Password     string    `json:"password"`
	RegisteredAt time.Time `json:"registeredAt"`
	Lists 		 []List    `json:"lists,omitempty"`
}

type UserInput struct {
	Name 		 string    `json:"name"`
	Email 		 string    `json:"email"`
	Phone 		 string    `json:"phone"`
	Password 	 string    `json:"password"`
}


