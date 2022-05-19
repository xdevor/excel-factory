package controller

import (
	"echoException/excel-faker/constvar"
	"echoException/excel-faker/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSourceOptions(c *gin.Context) {
	helper.ApiResponse(http.StatusOK, "ok", constvar.SourceOptions, c)
}

func GetDatTypeOptions(c *gin.Context) {
	helper.ApiResponse(http.StatusOK, "ok", constvar.DataTypeOptions, c)
}
