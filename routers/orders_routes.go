package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/gingormpostgres/controllers"
)

func configureOrdersRouter(router *gin.RouterGroup, controller *controllers.OrdersController) {
	router.GET("", controller.GetOrders)
	router.POST("", controller.CreateOrders)
	router.GET("/:id", func(ctx *gin.Context) {})
	router.PUT("/:id", func(ctx *gin.Context) {})
	router.DELETE("/:id", controller.DeleteOrders)
}
