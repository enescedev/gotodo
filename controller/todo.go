package controller

import (
	"github.com/enescedev/gotodo/config"
	"github.com/enescedev/gotodo/models"
	"github.com/gin-gonic/gin"
)

func GetToDo(c *gin.Context) {
	todos := []models.Todo{}
	config.DB.Find(&todos)
	c.JSON(200, &todos)
}

func CreateToDo(c *gin.Context) {
	var todos models.Todo
	c.BindJSON(&todos)
	config.DB.Create(&todos)
	c.JSON(200, &todos)
}

func DeleteToDo(c *gin.Context) {
	var todos models.Todo
	config.DB.Where("id = ?", c.Param("id")).Delete(&todos)
	c.JSON(200, &todos)

}

func UpdateToDo(c *gin.Context) {
	var todos models.Todo
	config.DB.Where("id = ?", c.Param("id")).First(&todos)
	c.BindJSON(&todos)
	config.DB.Save(&todos)
	c.JSON(200, &todos)

}
