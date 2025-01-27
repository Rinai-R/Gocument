package models

import "time"

// User 古希腊掌管用户信息的神
type User struct {
	Id        int64      `json:"id,omitempty" gorm:"primary_key;int;autoIncrement"`
	Username  string     `json:"username,omitempty" gorm:"unique;not null;varchar(50);index:idx_name_pass"`
	Password  string     `json:"password,omitempty" gorm:"not null;index:idx_name_pass"`
	Bio       string     `json:"bio,omitempty" gorm:"varchar(255);default:'这里填写个人简介'"`
	Gender    string     `json:"gender,omitempty" gorm:"varchar(50);default:'武装直升机'"`
	Avatar    string     `json:"avatar,omitempty" gorm:"varchar(255);default:'https://www.默认头像.com'"`
	Documents []Document `json:"documents,omitempty" gorm:"foreignKey:UserId;constraint;onDelete;CASCADE"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
