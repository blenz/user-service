package models

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"name"`
}
