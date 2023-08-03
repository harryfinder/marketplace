package models

import "time"

type User struct {
	ID             int64     `gorm:"primary_key"  json:"id"`
	FullName       string    `gorm:"size:255;not null"          json:"full_name"`
	Email          string    `gorm:"size:255;not null;unique"   json:"email"`
	Password       string    `gorm:"size:255;not null;"         json:"password"`
	NumberPhone    string    `gorm:"size:100;"                  json:"number_phone"`
	RepeatPassword string    `gorm:"size:255;not null;"         json:"repeat_password"`
	ActivePhone    bool      `gorm:"default:false"              json:"active_phone"`
	ActiveEmail    bool      `gorm:"default:false"              json:"active_email"`
	Role           int64     `gorm:"default:1"                  json:"role"`
	Status         bool      `gorm:"default:false"                   json:"stocks" `
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"createdAt"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"  json:"updatedAt"`
}
type ResponseUserCheckingKey struct {
	UserId int64 `json:"user_id"`
	Role   int64 `json:"role"`
}
type UserLogin struct {
	Login    string `gorm:"size:255;not null;unique"   json:"login"`    // В поля логин вводите email или номер телефон
	Password string `gorm:"size:100;not null;"         json:"password"` // Длина пароля не должно быть меньше
}
