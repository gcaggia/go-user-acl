package user

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CtrlDependencies struct {
	service Service
}

type controller struct {
	deps CtrlDependencies
}

func NewController() Controller {
	return &controller{
		deps: CtrlDependencies{
			service: NewService(),
		},
	}
}

func NewControllerWithDeps(deps CtrlDependencies) Controller {
	return &controller{
		deps: deps,
	}
}

func (ctrl *controller) GetAll(c *gin.Context) {
	c.JSON(200, ctrl.deps.service.GetAll())
}

func (ctrl *controller) GetByID(c *gin.Context) {
	user, err := ctrl.deps.service.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(404, nil)
	} else {
		c.JSON(200, user)
	}
}

func (ctrl *controller) Create(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user, err := ctrl.deps.service.Create(user)
	if err != nil {
		c.JSON(400, nil)
	} else {
		c.JSON(200, user)
	}
}

func (ctrl *controller) Update(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user, _ = ctrl.deps.service.Update(c.Param("id"), user)
	c.JSON(200, user)
}

func (ctrl *controller) Delete(c *gin.Context) {
	ctrl.deps.service.Delete(c.Param("id"))
	c.JSON(200, nil)
}
