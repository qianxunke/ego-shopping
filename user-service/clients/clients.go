package clients



import (
	"github.com/qianxunke/ego-shopping/ego-plugins/ego_redis"
	_ "github.com/qianxunke/ego-shopping/ego-plugins/ego_redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/qianxunke/ego-shopping/ego-plugins/db"
	 r "github.com/go-redis/redis"
)

var (
	redisClient *r.Client;
)


func Init()  {
	redisClient=ego_redis.Redis()
}

func GetRedis()(*r.Client)  {
	return redisClient
}
