package mw

import "github.com/gomodule/redigo/redis"

var addr = "127.0.0.1:16379"
var network = "tcp"
var Redis redis.Conn

func InitRedis() {
	var err error
	Redis, err = redis.Dial(network, addr)
	if err != nil {
		panic(err)
	}
}
