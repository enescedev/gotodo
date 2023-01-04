package routes

import "github.com/gin-gonic/gin"

func userRoute(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
}
