package models

type Create struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
