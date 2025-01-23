package models

import "time"

type Document struct {
	Id         int       `json:"id,omitempty" gorm:"primaryKey"`
	UserId     int       `json:"user_id,omitempty" gorm:"unique;not null"`
	Title      string    `json:"document_type,omitempty" gorm:"not null"`
	IsPrivate  bool      `json:"is_private,omitempty" gorm:"not null"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime;not null"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime;not null"`
}
