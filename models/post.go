package models

import "time"

type Post struct {
	Status      int32     `json:"status" db:"status"`
	CommunityId int64     `json:"community_id" db:"community_id"`
	Id          int64     `json:"id" db:"post_id"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
