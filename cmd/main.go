package main

import (
	"assigment-2-golang-faaza/internal/app"
	"assigment-2-golang-faaza/internal/domain"
	"assigment-2-golang-faaza/internal/infrastructure"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    dsn := "host=localhost user=postgres password=yourpassword dbname=yourdbname port=5432 sslmode=disable TimeZone=Asia/Jakarta"
    db, err := infrastructure.ConnectDB(dsn)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto Migrate
    db.AutoMigrate(&domain.Order{}, &domain.OrderItem{})

    // Initialize repositories and services
    orderRepo := app.NewOrderRepository(db)
    orderService := app.NewOrderService(orderRepo)

    // Initialize Gin
    r := gin.Default()

    // Routes
    r.POST("/orders", func(c *gin.Context) {
        var order domain.Order
        if err := c.BindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := orderService.CreateOrder(&order); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, order)
    })

    r.GET("/orders", func(c *gin.Context) {
        orders, err := orderService.GetOrders()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, orders)
    })

    r.PUT("/orders/:id", func(c *gin.Context) {
        var order domain.Order
        id := c.Param("id")
        if err := db.First(&order, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
            return
        }
        if err := c.BindJSON(&order); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := orderService.UpdateOrder(&order); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, order)
    })

	r.DELETE("/orders/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		if err := orderService.DeleteOrder(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
	})
	

    // Start server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
