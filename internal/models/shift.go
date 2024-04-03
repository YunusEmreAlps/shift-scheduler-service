package models

type Shift struct {
	ID    int    `json:"id"`
	Start string `json:"start"`
	End   string `json:"end"`
	User  User   `json:"user"`
}
