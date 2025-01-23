package models

import "time"

type User struct {
	Id        int64     `json:"id,omitempty" gorm:"primary_key;autoIncrement'"`
	Username  string    `json:"username,omitempty" gorm:"unique;not null;index:idx_name_pass"`
	Password  string    `json:"password,omitempty" gorm:"not null;index:idx_name_pass"`
	Bio       string    `json:"bio,omitempty" gorm:"default:'这里填写个人简介'"`
	Gender    string    `json:"gender,omitempty" gorm:"default:'武装直升机'"`
	Avatar    string    `json:"avatar,omitempty" gorm:"default:'https://www.默认头像.com'"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
