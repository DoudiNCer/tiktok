package handler

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/dal"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	dal.UserInfo(c)
}
