package models

import "time"

type Document struct {
	Id         int       `json:"omitempty" gorm:"primaryKey;autoIncrement"`
	Username   string    `json:"username,omitempty" gorm:"not null"`
	Title      string    `json:"title,omitempty" gorm:"not null"`
	IsPrivate  bool      `json:"is_private,omitempty" gorm:"not null"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime;not null"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime;not null"`
}

var EsDocument string

func init() {
	EsDocument = `{
    "mappings": {
        "properties": {
            "all": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "id": {
                "type": "keyword"
            },
            "title": {
                "type": "text",
                "analyzer": "ik_max_word",
                "copy_to": "all"
            },
            "content": {
                "type": "text",
                "analyzer": "ik_max_word",
                "copy_to": "all"
            }
        }
    }
}`

}
