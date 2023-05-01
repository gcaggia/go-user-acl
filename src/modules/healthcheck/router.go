package healthcheck

import "github.com/gin-gonic/gin"

type Router interface {
	SetupRouter(router *gin.Engine) *gin.Engine
}

type RouterDependencies struct {
	controller Controller
}

type router struct {
	deps RouterDependencies
}

func NewRouter() Router {
	return &router{
		deps: RouterDependencies{
			controller: NewController(),
		},
	}
}

func NewRouterWithDeps(deps RouterDependencies) Router {
	return &router{
		deps: deps,
	}
}

func (r *router) SetupRouter(router *gin.Engine) *gin.Engine {
	routes := router.Group("/healthcheck")
	{
		routes.GET("", func(c *gin.Context) { r.deps.controller.Ping(c) })
	}
	return router
}
