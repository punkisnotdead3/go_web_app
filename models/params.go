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
	// 1 点赞 -1 点踩 oneof 是限制这个值只能为多少
	Direction int64 `json:"direction,string" binding:"required,oneof=1 -1"`
}

type ParamListData struct {
	PageSize int64  `form:"pageSize" binding:"required"`
	PageNum  int64  `form:"pageNum" binding:"required"`
	Order    string `form:"order" binding:"required,oneof=time hot"`
}

const (
	DirectionLike   = 1
	DirectionUnLike = -1
	// 按照帖子时间排序
	OrderByTime = "time"
	// 按照点赞数量排序
	OrderByHot = "hot"
)
