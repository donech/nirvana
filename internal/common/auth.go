package common

import (
	"context"

	"github.com/donech/tool/xlog"

	"github.com/donech/tool/xjwt"
)

func GetUserID(ctx context.Context) int64 {
	claims := xjwt.GetClaimsFromCtx(ctx)
	xlog.S(ctx).Infof("claims: %#v", claims)
	switch id := claims["id"].(type) {
	case float64:
		return int64(id)
	default:
		return 0
	}
}
