package controllers

import (
	"errors"
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

// CreateOrders godoc
// @Summary Create an order
// @Description Create an order with items
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body requests.Order true "Order"
// @Success 201 {string} string "Order Created"
// @Router /orders [post]
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

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders with items
// @Tags Orders
// @Produce json
// @Success 200 {array} models.Orders "Orders"
// @Router /orders [get]
func (o OrdersController) GetOrders(ctx *gin.Context) {
	var orders []models.Orders
	result := o.Db.Preload("Items").Find(&orders)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(200, orders)
}

// GetOrderById godoc
// @Summary Get an order by ID
// @Description Get an order by ID with items
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} models.Orders "Order"
// @Router /orders/{id} [get]
func (o OrdersController) GetOrderById(ctx *gin.Context) {
	var orders models.Orders
	id := ctx.Param("id")
	result := o.Db.Preload("Items").First(&orders, id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	ctx.JSON(200, orders)
}

func (o OrdersController) UpdateOrder(ctx *gin.Context) {
	var orderID string = ctx.Param("id")
	var updatedOrderRequest requests.Order

	if err := ctx.ShouldBindJSON(&updatedOrderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingOrder models.Orders
	result := o.Db.First(&existingOrder, orderID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order"})
		}
		return
	}

	existingOrder.CustomerName = updatedOrderRequest.CustomerName

	tx := o.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Save(&existingOrder).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	if err := tx.Delete(&models.Item{OrderId: existingOrder.OrderId}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated items"})
		return
	}

	var newItems []models.Item
	for _, item := range updatedOrderRequest.Items {
		newItems = append(newItems, models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     existingOrder.OrderId,
		})
	}

	for _, item := range newItems {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create associated item"})
			return
		}
	}

	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}

// DeleteOrders godoc
// @Summary Delete an order
// @Description Delete an order with items
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {string} string "Order deleted"
// @Router /orders/{id} [delete]
func (o OrdersController) DeleteOrders(ctx *gin.Context) {
	var order models.Orders
	var item models.Item
	id := ctx.Param("id")

	result := o.Db.First(&order, id)

	tx := o.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if result.Error == nil {
		if err := o.Db.Where(&models.Item{OrderId: order.OrderId}).Delete(&item).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete associated items"})
			return
		}
	}

	if err := tx.Delete(&order).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
