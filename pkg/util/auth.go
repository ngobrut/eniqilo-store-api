package util

import (
	"context"

	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
)

func GetUserIDFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if userID, ok := ctx.Value(constant.UserIDKey).(string); ok {
		return userID
	}

	return ""
}
