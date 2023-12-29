package models

type Users struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password []byte `json:""`
	Email    string `json:"email" gorm:"unique"`
}
