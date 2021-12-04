package redis

import (
	"go_web_app/utils"

	"github.com/go-redis/redis"

	"go.uber.org/zap"
)

func getRedisKeyForLikeUserSet(postId int64) string {
	key := KeyPostLikeZetPrefix + utils.Int64ToString(postId)
	zap.L().Debug("getRedisKeyForLikeUserSet", zap.String("setKey", key))
	return key
}

// CheckLike 判断之前有没有投过票 true 代表之前 投过 false 代表之前没有投过
func CheckLike(postId int64, userId int64) (int64, bool) {
	like := rdb.ZScore(getRedisKeyForLikeUserSet(postId), utils.Int64ToString(userId))
	result, err := like.Result()
	if err != nil {
		zap.L().Error("checkLike error", zap.Error(err))
		return 0, false
	}
	zap.L().Info("checkLike val", zap.Float64(utils.Int64ToString(userId), like.Val()))
	return int64(result), true
}

// 点赞
func DoLike(postId int64, userId int64, direction int64) error {
	value := redis.Z{
		Score:  float64(direction),
		Member: utils.Int64ToString(userId),
	}
	_, err := rdb.ZAdd(getRedisKeyForLikeUserSet(postId), value).Result()
	if err != nil {
		zap.L().Error("doLike error", zap.Error(err))
		return err
	}
	return nil
}

func AddLike(postId int64, direction int64) error {
	_, err := rdb.ZIncrBy(KeyLikeNumberZSet, float64(direction), utils.Int64ToString(postId)).Result()
	if err != nil {
		zap.L().Error("AddLike error", zap.Error(err))
		return err
	}
	return nil
}
