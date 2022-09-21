package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(30);unique;not null"`
	Password string `json:"password" gorm:"type:varchar(32);not null"`
	Role     string `json:"role" gorm:"type:varchar(10);not null"`
	Active   bool   `json:"active" gorm:"type:bool;default:false"`
}
