package web

import (
	"gitee.com/nichanghao/gdmin/global"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer(engine *gin.Engine) {

	err := engine.Run(global.Config.Server.Address)
	if err != nil {
		log.Fatalf("gin run error: %v", err)
	}
}
