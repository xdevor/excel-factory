package controller

import (
	"echoException/excel-faker/dto"
	"echoException/excel-faker/helper"
	"echoException/excel-faker/service"

	"github.com/gin-gonic/gin"
)

func ExportFakeExcelFile(c *gin.Context) {
	input, isValid := parseRequestPayload(c)
	if !isValid || !validateExportFakeExcel(input) {
		apiReturn(isValid, "", c)
		return
	}

	data := service.GenerateFakeExcelData(input)
	fileNameWithPath := service.GenerateExcelFileName(input.FileName)
	ok := service.SaveExcel(data, fileNameWithPath)
	resultFileUrl := helper.GetAppUrl(c) + "/" + fileNameWithPath

	apiReturn(ok, resultFileUrl, c)
}

func parseRequestPayload(c *gin.Context) (dto.ExportFakeExcel, bool) {
	var input dto.ExportFakeExcel
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, false
	}
	return input, true
}

func validateExportFakeExcel(data dto.ExportFakeExcel) bool {
	if data.RowCount < 1 || len(data.Columns) < 1 {
		return false
	}
	return true
}

func apiReturn(ok bool, url string, c *gin.Context) {
	if ok {
		c.JSON(200, gin.H{
			"message": "success",
			"data":    url,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "bad",
			"data":    "wrong request payload",
		})
	}
}
