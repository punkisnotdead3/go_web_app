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

func GetPostDetail(id int64) (apiPostDetail *models.ApiPostDetail, err error) {
	//先查帖子实体
	post, err := mysql.GetPostDetail(id)
	//再查 作者 名称
	username, err := mysql.GetUserNameById(post.AuthorId)
	if err != nil {
		zap.L().Warn("no author ")
		err = nil
	}
	//再查板块实体
	community, err := GetCommunityById(post.CommunityId)
	if err != nil {
		zap.L().Warn("no community ")
		err = nil
	}
	apiPostDetail = new(models.ApiPostDetail)
	apiPostDetail.AuthorName = username

	apiPostDetail.Community = community
	apiPostDetail.Post = post

	return apiPostDetail, err

}
