package domain

type List struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	Tasks		 []Task	   `json:"tasks"`
	UserId		 uint64    `json:"userId"`
}