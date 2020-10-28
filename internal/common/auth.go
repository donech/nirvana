package common

import (
	"context"
	"encoding/json"

	"github.com/donech/tool/xjwt"
)

func GetUserID(ctx context.Context) int64 {
	claims := xjwt.GetClaimsFromCtx(ctx)
	switch id := claims["id"].(type) {
	case float64:
		return int64(id)
	case json.Number:
		v, _ := id.Int64()
		return v
	}
	return 0
}
