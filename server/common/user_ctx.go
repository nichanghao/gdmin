package common

import (
	"context"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"github.com/gin-gonic/gin"
)

var (
	CTX = &userContext{}
)

const ClaimsKey = "claims"

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

	if claims := (*ctx).Value("claims"); claims != nil {
		id := claims.(*UserClaims).ID
		return id
	}

	return 0

}
