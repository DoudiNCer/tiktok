package mw

import "github.com/gomodule/redigo/redis"

const addr = "127.0.0.1:16379"
const network = "tcp"

var Redis redis.Conn

func InitRedis() {
	var err error
	Redis, err = redis.Dial(network, addr)
	if err != nil {
		panic(err)
	}
}
