package utils

import (
	"gitee.com/nichanghao/gdmin/model/common"
	"github.com/gin-gonic/gin"
)

var (
	CLAIMS = &claimsUtil{}
)

type claimsUtil struct {
}

func (*claimsUtil) GetUserClaims(c *gin.Context) (*common.UserClaims, error) {
	if claims, exists := c.Get("claims"); !exists {

		return nil, common.ErrTokenIllegal
	} else {
		return claims.(*common.UserClaims), nil
	}
}
