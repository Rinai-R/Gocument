package models

import "time"

// Document 古希腊掌管文档基本信息的神
type Document struct {
	Id        int       `json:"id,omitempty" gorm:"int;primaryKey;autoIncrement"`
	UserId    int       `json:"user_id,omitempty" gorm:"int;not null;index:idx_user_id"`
	Title     string    `json:"title,omitempty" gorm:"varchar(50)not null"`
	IsPrivate bool      `json:"is_private,omitempty" gorm:"not null"`
	CreateAt  time.Time `json:"create_at,omitempty" gorm:"autoCreateTime;not null"`
	UpdateAt  time.Time `json:"update_at,omitempty" gorm:"autoUpdateTime;not null"`
}

// SearchDocument 搜索文档请求
type SearchDocument struct {
	UserId  int64  `json:"user_id,omitempty"`
	Content string `json:"content,omitempty"`
}

// ElasticDocument 古希腊掌管文档标题及内容的神
// 存储在ES中，便于实现查询功能
type ElasticDocument struct {
	Id        string    `json:"id,omitempty"`
	UserId    int64     `json:"user_id,omitempty"`
	IsPrivate bool      `json:"is_private,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreateAt  time.Time `json:"create_at,omitempty"`
	UpdateAt  time.Time `json:"update_at,omitempty"`
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
            "user_id": {
                "type": "keyword"
            },
            "is_private": {
                "type": "boolean"
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
            },
            "created_at": {
                "type": "date",
                "format": "strict_date_optional_time||epoch_millis"
            },
            "updated_at": {
                "type": "date",
                "format": "strict_date_optional_time||epoch_millis"
            }
        }
    }
}
`

}
