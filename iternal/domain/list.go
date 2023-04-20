package domain

type List struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
}

type ListInput struct {
	Name         string   `json:"name"`
}