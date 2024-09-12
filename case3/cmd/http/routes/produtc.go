package routes

import (
	"github.com/elangreza14/qbit/case3/controller"
	"github.com/gin-gonic/gin"
)

func ProductRoute(route *gin.RouterGroup, ProductController *controller.ProductController) {
	ProductRoutes := route.Group("/products")
	ProductRoutes.GET("", ProductController.ProductList())
}
