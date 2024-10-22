package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

// IOption connector 配置
type IOption interface {
	DataSourceName() string // redis://<user>:<pass>@localhost:6379/<db>
	IsDebug() bool

	// TakePoolSize 套接字连接的最大数量。
	// 默认是每个可用 CPU（由 runtime.GOMAXPROCS 报告）10个连接。
	TakePoolSize() int

	// TakeMinIdleConns
	// 最小空闲连接数，当建立新连接较慢时很有用。
	TakeMinIdleConns() int

	// TakeMaxConnAge 客户端关闭连接的生命周期。
	// 默认是不关闭老化连接。
	TakeMaxConnAge() time.Duration

	// TakePoolTimeout 如果所有连接都忙，客户端等待连接的时间长度，然后返回错误。
	// 默认是 ReadTimeout + 1秒。
	TakePoolTimeout() time.Duration

	// TakeIdleTimeout 客户端关闭空闲连接后的时间长度。
	// 应该小于服务器的超时时间。
	// 默认是5分钟。-1 禁用空闲超时检查。
	TakeIdleTimeout() time.Duration

	// TakeIdleCheckFrequency 空闲连接清理器进行空闲检查的频率。
	// 默认是1分钟。-1 禁用空闲连接清理器，
	// 但如果设置了 IdleTimeout，客户端仍会丢弃空闲连接。
	TakeIdleCheckFrequency() time.Duration
}

// IConnector connector 接口
type IConnector interface {
	Connect() *redis.Client
}
