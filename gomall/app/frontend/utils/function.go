package utils

import (
	"context"
	"fmt"
)

func GetUserIdFromCtx(ctx context.Context) int32 {
	fmt.Println("Get user id")
	userId := ctx.Value(SessionUserId)
	fmt.Println(userId)
	if userId == nil {
		return 0
	}
	return userId.(int32)
}
