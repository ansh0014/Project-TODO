package model

type Todo struct {
	ID   int    `json:"ID"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}
type User struct{
	Username string `json:"username"`
	Password string  `json:"password"`
}
