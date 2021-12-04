package logic

import (
	"errors"
	"go_web_app/dao/redis"
	"go_web_app/models"
)

var ErrAleadyLike = errors.New("不能重复点赞")
var ErrAleadyUnLike = errors.New("不能重复点踩")

func PostLike(postData *models.ParamLikeData, userId int64) error {
	// 查询之前有没有点过赞
	direction, flag := redis.CheckLike(postData.PostId, userId)

	if flag {
		// 如果之前点过赞 则要判断 这次是否是重复点赞
		if direction == postData.Direction && direction == models.DirectionLike {
			return ErrAleadyLike
		}
		// 如果之前点过赞 则要判断 这次是否是重复 点踩
		if direction == postData.Direction && direction == models.DirectionUnLike {
			return ErrAleadyUnLike
		}
	}

	err := redis.DoLike(postData.PostId, userId, postData.Direction)
	if err != nil {
		return err
	}
	return nil
}
