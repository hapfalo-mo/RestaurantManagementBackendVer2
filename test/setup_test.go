package test

import (
	"RestuarantBackend/routers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routers.SetRoutesAPI(r)
	return r
}
