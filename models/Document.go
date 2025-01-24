package models

import "time"

type Document struct {
	Id         int       `json:"id,omitempty" gorm:"int;primaryKey;autoIncrement"`
	UserId     int       `json:"user_id,omitempty" gorm:"int;not null;index:idx_user_id"`
	Title      string    `json:"title,omitempty" gorm:"varchar(50)not null"`
	IsPrivate  bool      `json:"is_private,omitempty" gorm:"not null"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime;not null"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime;not null"`
}

type ElasticDocument struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"user_id,omitempty"`
	Content string `json:"content,omitempty"`
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
