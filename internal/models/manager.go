package models

type Manager struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Mail        string `json:"mail"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
