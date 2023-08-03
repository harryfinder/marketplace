package models

type CheckEmail struct {
	ID     int64  `json:"id"`
	Key    string `json:"key"`
	UserId int64  `json:"user_id"`
	Email  string `json:"email"`
	Status bool   `json:"status" gorm:"default:true"`
}
