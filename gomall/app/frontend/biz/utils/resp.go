package utils

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	frontendUtils "github.com/readlnh/biz-demo/gomall/app/frontend/utils"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := ctx.Value(frontendUtils.SessionUserId)
	fmt.Printf("WarpResponse - context user_id type: %T, value: %v\n", userId, userId)
	content["user_id"] = frontendUtils.GetUserIdFromCtx(ctx)
	return content
}
