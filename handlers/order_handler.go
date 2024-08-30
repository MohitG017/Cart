package handlers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateOrderHandler handles converting a cart to an order
func CreateOrderHandler(c *gin.Context) {
	userToken := c.GetHeader("Authorization")
	if userToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	var user models.User
	if err := db.DB.Where("token = ?", userToken).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify if the cart belongs to the user
	var cart models.Cart
	if err := db.DB.Where("id = ? AND user_id = ?", order.CartID, user.ID).First(&cart).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Cart does not belong to user"})
		return
	}

	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// ListOrdersHandler handles listing all orders
func ListOrdersHandler(c *gin.Context) {
	var orders []models.Order
	if err := db.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
