package helper

import (
	"github.com/gin-gonic/gin"
)

func GetAppUrl(c *gin.Context) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host
}
