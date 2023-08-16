package router

import (
	"fmt"
	"github.com/WoodExplorer/user-auth/internal/configs"
	"github.com/WoodExplorer/user-auth/internal/middlewares"
	"github.com/WoodExplorer/user-auth/internal/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine  *gin.Engine
	roleSvc services.Role
	usrSvc  services.User
}

func InitRouter(
	roleSvc services.Role,
	usrSvc services.User,
) Router {
	var r Router

	r.roleSvc = roleSvc
	r.usrSvc = usrSvc

	engine := gin.New()
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(gin.Recovery())
	engine.Use(middlewares.GinLogger())
	r.engine = engine

	v1 := engine.Group("/api/v1")
	r.initUserRoutes(v1.Group("/users"))
	r.initRoleRoutes(v1.Group("/roles"))

	return r
}

func (r Router) initUserRoutes(g *gin.RouterGroup) {
	g.GET("", Wrapper(r.listUsers)) // TODO: 性能风险, 调试用
	g.GET("/:name", Wrapper(r.getUser))
	g.POST("/", Wrapper(r.createUser))
	g.DELETE("/:name", Wrapper(r.deleteUser))
}

func (r Router) initRoleRoutes(g *gin.RouterGroup) {
	g.GET("/:name", Wrapper(r.getRole))
	g.POST("/", Wrapper(r.createRole))
	g.DELETE("/:name", Wrapper(r.deleteRole))
}

func (r Router) Start() (err error) {
	return r.engine.Run(fmt.Sprintf(":%d", configs.GetPort()))
}
