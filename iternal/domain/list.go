package domain

type List struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Tasks		 []Task	   `json:"tasks"`
}