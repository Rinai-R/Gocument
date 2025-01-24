package models

// Permission 表示权限管理的表
// 在文档为私有状态时，可以单独为每个人赋予只读，或者读写权限
// 在文档为公开状态时，默认每个人都有只读权限，写入权限通过该表赋予。
// type字段表示权限类型，false表示只读，true表示读写。
type Permission struct {
	DocumentId int      `json:"document_id,omitempty" gorm:"not null;index:documentid_userid_idx"`
	UserId     int      `json:"user_id,omitempty" gorm:"not null;index:documentid_userid_idx"`
	Type       bool     `json:"type,omitempty" gorm:"not null"`
	Document   Document `gorm:"foreignKey:DocumentId"`
	User       User     `gorm:"foreignKey:UserId"`
}
