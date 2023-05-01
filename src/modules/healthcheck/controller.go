package healthcheck

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Ping(c *gin.Context)
}

type CtrlDependencies struct {
}

type controller struct {
	deps CtrlDependencies
}

func NewController() Controller {
	return &controller{}
}

func NewControllerWithDeps(deps CtrlDependencies) Controller {
	return &controller{
		deps: deps,
	}
}

func (ctrl *controller) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
