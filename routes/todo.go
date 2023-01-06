package routes

import (
	"github.com/enescedev/gotodo/controller"
	"github.com/gin-gonic/gin"
)

func ToDoRoute(router *gin.Engine) {
	router.GET("/", controller.GetToDo)
	router.POST("/", controller.CreateToDo)
	router.DELETE("/:id", controller.DeleteToDo)
	router.PUT("/:id", controller.UpdateToDo)
}
