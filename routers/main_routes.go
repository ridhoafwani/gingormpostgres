package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/gingormpostgres/config"
	"github.com/ridhoafwani/gingormpostgres/controllers"
	docs "github.com/ridhoafwani/gingormpostgres/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// @title Orders API
// @version 1.0
// @description This is an Orders API.

func NewMainRouter() *gin.Engine {
	dbName := "orders_by"
	db := config.DatabaseConnection(&dbName)
	ordersController := controllers.NewOrdersController(db)

	service := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome")
	})

	service.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not Found")
	})

	router := service.Group("/api/v1")
	ordersRouter := router.Group("/orders")

	configureOrdersRouter(ordersRouter, ordersController)

	return service
}
