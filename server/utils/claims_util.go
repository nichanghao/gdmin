package utils

import (
	"gitee.com/nichanghao/gdmin/common"
	"gitee.com/nichanghao/gdmin/common/buserr"
	"github.com/gin-gonic/gin"
)

var (
	CLAIMS = &claimsUtil{}
)

type claimsUtil struct {
}

func (*claimsUtil) GetUserClaims(c *gin.Context) (*common.UserClaims, error) {
	if claims, exists := c.Get("claims"); !exists {

		return nil, buserr.NewTokenAuthErr("invalid token")
	} else {
		return claims.(*common.UserClaims), nil
	}
}
