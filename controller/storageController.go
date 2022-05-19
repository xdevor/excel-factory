package controller

import "github.com/gin-gonic/gin"

func GetPublicFile(c *gin.Context) {
	folder := c.Param("folder")
	fileName := c.Param("file")
	path := "storage/public/" + folder + "/" + fileName
	c.File(path)
}
