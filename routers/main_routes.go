package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/gingormpostgres/config"
	"github.com/ridhoafwani/gingormpostgres/controllers"
)

func NewMainRouter() *gin.Engine {
	dbName := "orders_by"
	db := config.DatabaseConnection(&dbName)
	ordersController := controllers.NewOrdersController(db)

	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not Found")
	})

	router := service.Group("/api/v1")
	ordersRouter := router.Group("/orders")

	configureOrdersRouter(ordersRouter, ordersController)

	return service
}
