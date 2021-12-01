package logic

import (
	"go_web_app/dao/mysql"
	"go_web_app/models"
	"go_web_app/pkg/snowflake"
	"strconv"

	"go.uber.org/zap"
)

// chuan
func CreatePost(post *models.Post) (msg string, err error) {
	// 雪花算法 生成帖子id
	post.Id = snowflake.GenId()
	zap.L().Debug("createPostLogic", zap.Int64("postId", post.Id))
	err = mysql.InsertPost(post)
	if err != nil {
		return "failed", err
	}
	//发表帖子成功时 要把帖子id 回给 请求方
	return strconv.FormatInt(post.Id, 10), nil
}
