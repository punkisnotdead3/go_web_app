package models

//定义请求的参数结构体

type ParamRegister struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	// eqfield 指定必须相等的字段
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamPost struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CommunityId int64  `json:"community_id" binding:"required"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamLikeData struct {
	// 帖子id
	PostId int64 `json:"post_id,string" binding:"required"`
	// 1 点赞 -1 点踩
	Direction int64 `json:"direction,string" binding:"required,oneof=1 -1"`
}

const (
	DirectionLike   = 1
	DirectionUnLike = -1
)
