package logic

import (
	"go_web_app/dao/mysql"
	"go_web_app/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()

}
