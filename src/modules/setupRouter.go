package modules

import (
	HealthcheckModule "github.com/gcaggia/go-user-acl/src/modules/healthcheck"
	UserModule "github.com/gcaggia/go-user-acl/src/modules/user"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	UserModule.NewRouter().SetupRouter(router)
	HealthcheckModule.NewRouter().SetupRouter(router)

	return router
}
