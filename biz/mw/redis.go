package mw

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

const addr = "127.0.0.1:16379"
const network = "tcp"

var Redis redis.Conn
var RedisPoll *redis.Pool

func InitRedis() {
	RedisPoll = &redis.Pool{
		// 连接方法
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, addr,
				redis.DialReadTimeout(1*time.Second),
				redis.DialWriteTimeout(1*time.Second),
				redis.DialConnectTimeout(1*time.Second),
			)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		// 最大的空闲连接数，
		// 表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: 256,
		// 最大的激活连接数，表示同时最多有N个连接
		MaxActive: 256,
		// 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: time.Duration(120),
	}
}
