package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/readlnh/biz-demo/gomall/app/frontend/utils"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		fmt.Printf("GlobalAuth - session user_id type: %T, value: %v\n", userId, userId)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, userId)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		if userId == nil {
			c.Redirect(consts.StatusFound, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
