package routes

import (
	"github.com/elangreza14/qbit/case3/controller"
	"github.com/elangreza14/qbit/case3/middleware"
	"github.com/gin-gonic/gin"
)

func CartRoute(route *gin.RouterGroup, CartController *controller.CartController, authMiddleware *middleware.AuthMiddleware) {
	CartRoutes := route.Group("/carts")
	CartRoutes.POST("", authMiddleware.MustAuthMiddleware(), CartController.AddProductToCartList())
	CartRoutes.GET("", authMiddleware.MustAuthMiddleware(), CartController.CheckAvailabilityCartList())
}
