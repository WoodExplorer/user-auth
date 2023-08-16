package router

import (
	"fmt"
	"github.com/WoodExplorer/user-auth/internal/configs"
	"github.com/WoodExplorer/user-auth/internal/middlewares"
	"github.com/WoodExplorer/user-auth/internal/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	usrSvc services.User
}

func InitRouter(usrSvc services.User) Router {
	var r Router

	r.usrSvc = usrSvc

	engine := gin.New()
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(gin.Recovery())
	engine.Use(middlewares.GinLogger())
	r.engine = engine

	v1 := engine.Group("/api/v1")
	r.initUserRoutes(v1.Group("/users"))

	return r
}

func (r Router) Start() (err error) {
	return r.engine.Run(fmt.Sprintf(":%d", configs.GetPort()))
}

func (r Router) initUserRoutes(g *gin.RouterGroup) {
	g.GET("/:name", Wrapper(r.getUser))
	g.POST("/", Wrapper(r.createUser))
	g.DELETE("/:name", Wrapper(r.deleteUser))
}
