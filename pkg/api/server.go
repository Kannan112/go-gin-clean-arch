package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kannan112/go-gin-clean-arch/cmd/api/docs"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/api/handler/interface"
)

type ServerHTTP struct {
	engine *gin.Engine
	port   string
}

func NewServerHTTP(userHandler interfaces.UserHandler) *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(sh.port)
}
