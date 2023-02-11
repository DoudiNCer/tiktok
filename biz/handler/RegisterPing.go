package handler

import (
	"context"
	"github.com/DodiNCer/tiktok/Controller"
	"github.com/cloudwego/hertz/pkg/app"
)

func RegisterPing(ctx context.Context, c *app.RequestContext) {
	Controller.Register(c)
}
