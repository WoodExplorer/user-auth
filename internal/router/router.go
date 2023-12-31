package router

import (
	"fmt"
	"github.com/WoodExplorer/user-auth/internal/configs"
	"github.com/WoodExplorer/user-auth/internal/middlewares"
	"github.com/WoodExplorer/user-auth/internal/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine     *gin.Engine
	roleSvc    services.Role
	usrSvc     services.User
	usrRoleSvc services.UserRole
	authnSvc   services.Authn
	authzSvc   services.Authz
}

func InitRouter(
	roleSvc services.Role,
	usrSvc services.User,
	usrRoleSvc services.UserRole,
	authnSvc services.Authn,
	authzSvc services.Authz,
) (engine *gin.Engine, r Router) {

	r.roleSvc = roleSvc
	r.usrSvc = usrSvc
	r.usrRoleSvc = usrRoleSvc
	r.authnSvc = authnSvc
	r.authzSvc = authzSvc

	engine = gin.New()
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(gin.Recovery())
	engine.Use(middlewares.GinLogger())
	r.engine = engine

	v1 := engine.Group("/api/v1")
	r.initUserRoutes(v1.Group("/users"))
	r.initRoleRoutes(v1.Group("/roles"))
	r.initUserRoleRoutes(v1.Group("/user-roles"))
	r.initAuthnRoutes(v1.Group("/authn"))
	r.initAuthzRoutes(v1.Group("/authz"))

	return
}

func (r Router) initUserRoutes(g *gin.RouterGroup) {
	g.GET("", Wrapper(r.listUsers)) // TODO: 性能风险, 仅调试用
	g.GET("/:name", Wrapper(r.getUser))
	g.POST("/", Wrapper(r.createUser))
	g.DELETE("/:name", Wrapper(r.deleteUser))
}

func (r Router) initRoleRoutes(g *gin.RouterGroup) {
	g.GET("", Wrapper(r.listRoles)) // TODO: 性能风险, 仅调试用
	g.GET("/:name", Wrapper(r.getRole))
	g.POST("/", Wrapper(r.createRole))
	g.DELETE("/:name", Wrapper(r.deleteRole))
}

func (r Router) initUserRoleRoutes(g *gin.RouterGroup) {
	g.POST("/", Wrapper(r.bindUserRole))
}

func (r Router) initAuthnRoutes(g *gin.RouterGroup) {
	g.POST("/tokens", Wrapper(r.applyToken))
	g.DELETE("/tokens", Wrapper(r.invalidateToken))
}

func (r Router) initAuthzRoutes(g *gin.RouterGroup) {
	g.GET("/check-role", Wrapper(r.checkRole))
	g.GET("/user-roles", Wrapper(r.getUserRoles))
}

func (r Router) Start() (err error) {
	return r.engine.Run(fmt.Sprintf(":%d", configs.GetPort()))
}

func (r Router) Stop() {
	return
}
