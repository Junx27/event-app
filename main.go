package main

import (
	"github.com/Junx27/event-app/config"
	"github.com/Junx27/event-app/database"
	"github.com/Junx27/event-app/repository"
	"github.com/Junx27/event-app/router"
	"github.com/Junx27/event-app/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewEnvConfig()
	db := database.Init(cfg, database.DBMigrator)
	r := gin.Default()
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	router.SetupAuthRouter(r, authService.(*service.AuthService))
	router.SetupUserRouter(r, db)
	router.SetupEventRouter(r, db)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
