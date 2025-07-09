package model

type Todo struct {
	ID   int    `json:"ID"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}
