package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type connector struct {
	opt IOption
}

func NewConnector(opt IOption) IConnector {
	return &connector{opt: opt}
}

// todo:  暂时其他配置用不到, 有时间再说

func (m connector) Connect() (db *redis.Client) {
	var (
		err error
		opt *redis.Options
	)

	if opt, err = redis.ParseURL(m.opt.DataSourceName()); err != nil {
		panic(err)
	}

	db = redis.NewClient(opt)
	if err = db.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return
}
