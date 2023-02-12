package handler

import (
	"context"
	"github.com/DodiNCer/tiktok/biz/handler/user_gorm"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	user_gorm.UserInfo(c)
}
