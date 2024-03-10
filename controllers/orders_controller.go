package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ridhoafwani/gingormpostgres/models"
	"github.com/ridhoafwani/gingormpostgres/requests"
	"gorm.io/gorm"
)

type OrdersController struct {
	Db *gorm.DB
}

func NewOrdersController(db *gorm.DB) *OrdersController {
	return &OrdersController{Db: db}
}

func (o OrdersController) CreateOrders(ctx *gin.Context) {
	var orderRequest requests.Order
	if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	order := models.Orders{
		CustomerName: orderRequest.CustomerName,
		OrderedAt:    time.Now(),
	}

	o.Db.Create(&order)

	var items []models.Item

	for _, item := range orderRequest.Items {
		items = append(items, models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     order.OrderId,
		})
	}

	for _, item := range items {
		o.Db.Create(&item)
	}
	ctx.JSON(http.StatusCreated, "Order Created")
}

func (o OrdersController) GetOrders(ctx *gin.Context) {
	var orders []models.Orders
	// Include related Items using Preload
	result := o.Db.Preload("Items").Find(&orders)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(200, orders)
}

// func (o OrdersController) GetOrders(ctx *gin.Context) {
// 	var orders []models.Orders
// 	o.Db.Find(&orders)

// 	for order := range orders {
// 		o.Db.Model(&orders[order]).Association("Items").Find(&orders[order].Items)
// 	}

// 	ctx.JSON(200, orders)
// }

func (o OrdersController) GetOrdersById(ctx *gin.Context) {
	var orders models.Orders
	id := ctx.Param("id")
	o.Db.First(&orders, id)
	ctx.JSON(200, orders)
}

func (o OrdersController) UpdateOrders(ctx *gin.Context) {
	var orders models.Orders
	id := ctx.Param("id")
	o.Db.First(&orders, id)
	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	o.Db.Save(&orders)
	ctx.JSON(200, orders)
}

func (o OrdersController) DeleteOrders(ctx *gin.Context) {
	var orders models.Orders
	id := ctx.Param("id")
	o.Db.First(&orders, id)
	o.Db.Delete(&orders)
	ctx.JSON(200, gin.H{"message": "Order deleted"})

}