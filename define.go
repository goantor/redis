package redis

import "github.com/go-redis/redis/v8"

// IOption connector 配置
type IOption interface {
	DataSourceName() string // redis://<user>:<pass>@localhost:6379/<db>
	IsDebug() bool
}

// IConnector connector 接口
type IConnector interface {
	Connect() *redis.Client
}
