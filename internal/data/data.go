package data

import (
	"context"
	"fmt"
	"wuzigoweb/internal/conf"
	"wuzigoweb/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo, NewPostRepo, NewAppRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	// 嵌入一个 grpc 的 client 端，通过这个 client 来调用我的 rpc 服务
	// rpc v1.UserClient

	log *log.Helper
	// 嵌入一个 gen 框架和 Redis 对象
	query     *query.Query
	redisConn *redis.Client
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger, red *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// 为 gen 生成的 query 代码设置数据库对象
	query.SetDefault(db)
	return &Data{query: query.Q, log: log.NewHelper(logger), redisConn: red}, cleanup, nil
}

// NewDB 用 gorm 连接数据库
func NewDB(cfg *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.Database.GetSource()))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	return db, err
}

// NewRedis 连接 Redis
func NewRedis(cfg *conf.Data) (*redis.Client, error) {
	redisConn := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.PassWord,
		DB:       int(cfg.Redis.Db),
		PoolSize: int(cfg.Redis.PoolSize),
	})
	if redisConn == nil {
		panic("fail to connect redis")
	}

	_, err := redisConn.Set(context.Background(), "abc", 100, time.Millisecond*60).Result()

	// 调用 redisConn.Ping() 方法来测试与 Redis 的连接是否正常
	_, err = redisConn.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to ping redis")
	}

	return redisConn, err
}
