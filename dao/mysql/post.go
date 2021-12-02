package mysql

import (
	"database/sql"
	"go_web_app/models"

	"go.uber.org/zap"
)

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
