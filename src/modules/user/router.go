package user

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
	routes := router.Group("/user")
	{
		routes.GET("", func(c *gin.Context) { r.deps.controller.GetAll(c) })
		routes.GET("/:id", func(c *gin.Context) { r.deps.controller.GetByID(c) })
		routes.POST("", func(c *gin.Context) { r.deps.controller.Create(c) })
		routes.PUT("/:id", func(c *gin.Context) { r.deps.controller.Update(c) })
		routes.DELETE("/:id", func(c *gin.Context) { r.deps.controller.Delete(c) })
	}
	return router
}
