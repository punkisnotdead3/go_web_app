package mysql

import (
	"database/sql"
	"go_web_app/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name,create_time from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if err == sql.ErrNoRows {
			zap.L().Warn("no community ")
			err = nil
		}
	}
	return

}

func GetCommunityById(id int64) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := "select community_id,community_name,introduction,create_time " +
		"from community where community_id=?"
	err = db.Get(community, sqlStr, id)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if err == sql.ErrNoRows {
			zap.L().Warn("no community ")
			err = nil
		}
	}
	return community, err

}
