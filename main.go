package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"net/http"

	"github.com/arsalanaa44/basketball/controllers"
	"github.com/arsalanaa44/basketball/initializers"
	"github.com/arsalanaa44/basketball/routes"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	BasketController      controllers.BasketController
	BasketRouteController routes.BasketRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	BasketController = controllers.NewBasketController(initializers.DB)
	BasketRouteController = routes.NewRouteBasketController(BasketController)

	server = gin.Default()
}

func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:" + config.ServerPort, config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	BasketRouteController.BasketRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
