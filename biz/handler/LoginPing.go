package handler

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/dal"
	"github.com/cloudwego/hertz/pkg/app"
)

func LoginPing(ctx context.Context, c *app.RequestContext) {
	dal.Login(c)
}
