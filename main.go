package main

import (
	"github.com/Junx27/ticket-booking/router"
	"github.com/gin-gonic/gin"
	"guthub.com/Junx27/event-app/config"
	"guthub.com/Junx27/event-app/database"
)

func main() {
	cfg := config.NewEnvConfig()
	db := database.Init(cfg, database.DBMigrator)
	r := gin.Default()
	router.SetupUserRouter(r, db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
