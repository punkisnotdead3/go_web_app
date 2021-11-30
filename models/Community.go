package models

import "time"

//omitempty 注意这个含义 如果为空的话可以不展示这个字段
type Community struct {
	Id          int64  `json:"id" db:"community_id"`
	Name        string `json:"name" db:"community_name"`
	Introdution string `json:"introdution,omitempty" db:"introduction" `
	// 也可以用int64 来直接返回时间戳
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
