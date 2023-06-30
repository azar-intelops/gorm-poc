package models

type User struct {
	Id        int64   `gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}