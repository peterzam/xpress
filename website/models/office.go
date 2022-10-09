package models

type Office struct {
	Id           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"type:varchar(30);not null"`
	Location_lat string `json:"location_lat" gorm:"type:varchar(10);not null"`
	Location_lng string `json:"location_lng" gorm:"type:varchar(10);not null"`
	Address      string `json:"address" gorm:"type:varchar(50);not null"`
	Phone        string `json:"phone" gorm:"type:varchar(20);not null"`
}
