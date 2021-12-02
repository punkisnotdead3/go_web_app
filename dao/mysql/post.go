package mysql

import (
	"database/sql"
	"go_web_app/models"
	"strconv"

	"go.uber.org/zap"
)

func GetPostList(offset int64, pageSize int64) (posts []*models.Post, err error) {
	zap.L().Info("GetPostList", zap.String("offset", strconv.FormatInt(offset, 10)), zap.String("pageSize", strconv.FormatInt(pageSize, 10)))
	sqlStr := "select post_id,title,content,author_id,community_id,create_time,update_time " +
		" from post limit ?,?"
	posts = make([]*models.Post, 0, pageSize)
	err = db.Select(&posts, sqlStr, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return posts, err

}

func GetPostDetail(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := "select post_id,title,content,author_id,community_id,create_time,update_time " +
		" from post where post_id=?"
	err = db.Get(post, sqlStr, id)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if err == sql.ErrNoRows {
			zap.L().Warn("no community ")
			err = nil
		}
	}
	return post, err
}

func InsertPost(post *models.Post) error {

	sqlstr := `insert into post(post_id,title,content,author_id,community_id) values(?,?,?,?,?)`
	_, err := db.Exec(sqlstr, post.Id, post.Title, post.Content, post.AuthorId, post.CommunityId)
	if err != nil {
		zap.L().Error("InsertPost dn error", zap.Error(err))
		return err
	}
	return nil
}
