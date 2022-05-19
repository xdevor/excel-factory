package router

import (
	"echoException/excel-faker/controller"

	"github.com/gin-gonic/gin"
)

func setStaticFile(r *gin.Engine) {
	r.GET("/", controller.HTMLhandler)
	r.GET("/:page", controller.HTMLhandler)
}
