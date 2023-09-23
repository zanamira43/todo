package models

// creating model for date
type TodoDate struct {
	ID   uint   `json:"id"`
	Date string `json:"date"`
	Todo []Todo `json:"-"`
}
