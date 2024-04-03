package models

type Organization struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Mail        string `json:"mail"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
