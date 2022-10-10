package models

type Package struct {
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Code       string `json:"code" gorm:"type:varchar(10);unique;not null"`
	Dest_name  string `json:"dest_name" gorm:"type:varchar(30);not null"`
	Dest_addr  string `json:"dest_addr" gorm:"type:varchar(50);not null"`
	Dest_phone string `json:"dest_phone" gorm:"type:varchar(20);not null"`
	Size       string `json:"size" gorm:"type:varchar(50);not null"`
	Type       string `json:"type" gorm:"type:varchar(20);default:normal;not null"`
	Note       string `json:"note" gorm:"type:varchar(50)"`
	Active     bool   `json:"active" gorm:"type:bool;default:false"`
	Src_id     uint   `json:"src_id" gorm:"type:uint;not null"`
	Created_at int64  `json:"created_at" gorm:"autoCreateTime"`
}
