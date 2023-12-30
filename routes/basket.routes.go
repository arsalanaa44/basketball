package routes

import (
	"github.com/arsalanaa44/basketball/controllers"
	"github.com/arsalanaa44/basketball/middleware"
	"github.com/gin-gonic/gin"
)

type BasketRouteController struct {
	basketController controllers.BasketController
}

func NewRouteBasketController(postController controllers.BasketController) BasketRouteController {
	return BasketRouteController{postController}
}

func (bc *BasketRouteController) BasketRoute(rg *gin.RouterGroup) {

	router := rg.Group("baskets")
	router.Use(middleware.DeserializeUser())
	router.POST("/", bc.basketController.CreateBasket)
	router.GET("/", bc.basketController.FindBaskets)
	router.PUT("/:basketId", bc.basketController.UpdateBasket)
	router.GET("/:basketId", bc.basketController.FindBasketById)
	router.DELETE("/:basketId", bc.basketController.DeleteBasket)
}
