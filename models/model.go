package models

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"primary_key;autoIncrement'"`
	Username  string    `json:"username" gorm:"unique;not null;index:idx_name_pass"`
	Password  string    `json:"password" gorm:"not null;index:idx_name_pass"`
	Bio       string    `json:"bio" gorm:"default:'这里填写个人简介'"`
	Gender    string    `json:"gender" gorm:"default:'武装直升机'"`
	Avatar    string    `json:"avatar" gorm:"default:'https://www.默认头像.com'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
