package main

import (
	"github.com/gin-gonic/gin"
	"guthub.com/Junx27/event-app/config"
	"guthub.com/Junx27/event-app/database"
	"guthub.com/Junx27/event-app/repository"
	"guthub.com/Junx27/event-app/router"
	"guthub.com/Junx27/event-app/service"
)

func main() {
	cfg := config.NewEnvConfig()
	db := database.Init(cfg, database.DBMigrator)
	r := gin.Default()
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	router.SetupAuthRouter(r, authService.(*service.AuthService))
	router.SetupUserRouter(r, db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
