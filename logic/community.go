package logic

import (
	"go_web_app/dao/mysql"
	"go_web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()

}

func GetCommunityById(id int64) (model *models.Community, err error) {
	return mysql.GetCommunityById(id)

}
