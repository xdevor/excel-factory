package router

import (
	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.Default()

	r.Delims("{[{", "}]}")
	r.LoadHTMLGlob("resource/html/*")

	setStaticFile(r)
	setApiRoute(r)
	setStorageRoute(r)

	return r
}
