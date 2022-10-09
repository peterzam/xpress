package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Phone    string `json:"phone" gorm:"type:varchar(20);unique;not null"`
	Name     string `json:"name" gorm:"type:varchar(30);not null"`
	Password string `json:"-" gorm:"type:varchar(32);not null"`
	Address  string `json:"address" gorm:"type:varchar(50);not null"`
	Role     string `json:"role" gorm:"type:varchar(10);default:user;not null"`
}
