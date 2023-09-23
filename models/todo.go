package models

// creating model for todo
type Todo struct {
	Id         uint     `json:"id"`
	Title      string   `json:"title"`
	Duration   float32  `json:"duration"`
	Completed  bool     `json:"completed"`
	TodoDateID uint     `json:"date_id"`
	TodoDate   TodoDate `json:"todo_date" gorm:"foreignKey:TodoDateID"`
}
