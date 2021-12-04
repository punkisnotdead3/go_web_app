package redis

// redis 集群可能在公司内部是公用的，所以定义key的时候 要加一些属于自己项目的前缀
const (
	// KeyPostTimeZSet key是 帖子id value 是帖子的发布时间
	KeyPostTimeZSet = "bbs:post:time"
	// KeyLikeNumberZSet key是 帖子id value 是帖子的点赞数量
	KeyLikeNumberZSet = "bbs:post:like:number"
	// KeyPostLikeZetPrefix  key是userid value是点赞或者点踩 后面需要拼接帖子的id
	KeyPostLikeZetPrefix = "bbs:post:like:postId:"
)
