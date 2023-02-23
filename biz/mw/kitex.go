package mw

import (
	"github.com/DodiNCer/tiktok/biz/mw/kitex_ffmpeg/kitexffmpeg"
	client "github.com/cloudwego/kitex/client"
)

var RPCClient kitexffmpeg.Client

const (
	DESTSERVICE = "KitexFfmpeg"
	HOSTPORT    = "0.0.0.0:19427"
)

func InitRPC() {
	rpcClient, err := kitexffmpeg.NewClient(DESTSERVICE, client.WithHostPorts(HOSTPORT))
	if err != nil {
		panic("RPC Connect Failed")
		return
	}
	RPCClient = rpcClient
}
