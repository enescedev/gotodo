package main

import (
	"log"

	"github.com/enescedev/gotodo/config"
	"github.com/enescedev/gotodo/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var cache *redis.Client

func main() {

	// init cache connection
	cache = config.CacheConnection("localhost:6379", "")
	log.Println("Redis e bağlandı")
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	config.Connect("postgres://bbpguser:bbpguserpassword@localhost:5432/postgres")
	routes.ToDoRoute(router)
	log.Println("Sunucu ayağa kalktı")

	//port
	router.Run(":8080")
}
