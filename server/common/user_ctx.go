package common

import (
	"context"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"github.com/gin-gonic/gin"
)

var (
	USER_CTX = &userContext{}
)

const (
	ClaimsKey = "claims"

	RequestKey = "_request"
)

type userContext struct {
}

func (*userContext) GetUserClaims(c *gin.Context) (*UserClaims, error) {
	if claims, exists := c.Get(ClaimsKey); !exists {

		return nil, buserr.NewTokenAuthErr("invalid token")
	} else {
		return claims.(*UserClaims), nil
	}
}

func (*userContext) GetUserId(ctx *context.Context) uint64 {

	if claims := (*ctx).Value(ClaimsKey); claims != nil {
		id := claims.(*UserClaims).ID
		return id
	}

	return 0

}
