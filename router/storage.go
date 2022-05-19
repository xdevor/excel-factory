package router

import (
	"echoException/excel-faker/controller"

	"github.com/gin-gonic/gin"
)

func setStorageRoute(r *gin.Engine) {
	r.GET("/storage/public/:folder/:file", controller.GetPublicFile)
}
