package main

import (
	"github.com/gin-gonic/gin"
	"./routes"
)

func main() {

	router := gin.New()
	routes.userRoute(router)
	router.Run(":8080")
}
