package models

type Office struct {
	Id           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"type:varchar(30);not null"`
	Location_lat string `json:"location_lat" gorm:"type:varchar(10);not null"`
	Location_lng string `json:"location_lng" gorm:"type:varchar(10);not null"`
	Address      string `json:"address" gorm:"type:varchar(50);not null"`
	Phone        string `json:"phone" gorm:"type:varchar(20);not null"`
}

type Package struct {
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Code       string `json:"code" gorm:"type:varchar(10);unique;not null"`
	Dest_name  string `json:"dest_name" gorm:"type:varchar(30);not null"`
	Dest_addr  string `json:"dest_addr" gorm:"type:varchar(50);not null"`
	Dest_phone string `json:"dest_phone" gorm:"type:varchar(20);not null"`
	Size       string `json:"size" gorm:"type:varchar(50);not null"`
	Type       string `json:"type" gorm:"type:varchar(20);default:normal;not null"`
	Note       string `json:"note" gorm:"type:varchar(50)"`
	Status     string `json:"status" gorm:"type:bool;default:false"`
	Src_id     uint   `json:"src_id" gorm:"type:uint;not null"`
	Created_at int64  `json:"created_at" gorm:"autoCreateTime"`
}

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Phone    string `json:"phone" gorm:"type:varchar(20);unique;not null"`
	Name     string `json:"name" gorm:"type:varchar(30);not null"`
	Password string `json:"-" gorm:"type:varchar(32);not null"`
	Address  string `json:"address" gorm:"type:varchar(50);not null"`
	Role     string `json:"role" gorm:"type:varchar(10);default:user;not null"`
}

type Chart struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

type Pickup struct {
	Id     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Size   string `json:"size" gorm:"type:varchar(50);not null"`
	Type   string `json:"type" gorm:"type:varchar(20);default:normal;not null"`
	Note   string `json:"note" gorm:"type:varchar(50)"`
	Src_id uint   `json:"src_id" gorm:"type:uint;not null"`
}

type Complaint struct {
	Id     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Note   string `json:"note" gorm:"type:varchar(50)"`
	Src_id uint   `json:"src_id" gorm:"type:uint;not null"`
}
