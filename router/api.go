package router

import (
	"echoException/excel-faker/controller"

	"github.com/gin-gonic/gin"
)

func setApiRoute(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/excel", controller.ExportFakeExcelFile)
		api.GET("/excel/options/source", controller.GetSourceOptions)
		api.GET("/excel/options/data-types", controller.GetDatTypeOptions)
	}
}
