package domain

type List struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	Tasks		 []Task	  `json:"tasks"`
	UserId		 int64    `json:"userId"`
}