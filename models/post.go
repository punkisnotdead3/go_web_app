package models

import "time"

type Post struct {
	Status      int32     `json:"status" db:"status"`
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"`
	Id          int64     `json:"id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	Title       string    `json:"title" db:"title" binding:"required" `
	Content     string    `json:"content" db:"content" binding:"required" `
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Community `json:"_community"`
	*Post      `json:"_post"`
}
