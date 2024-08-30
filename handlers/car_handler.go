package handlers

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCartHandler handles creating a new cart and adding items
func CreateCartHandler(c *gin.Context) {
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

	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart.UserID = user.ID
	if err := db.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	c.JSON(http.StatusCreated, cart)
}

// ListCartsHandler handles listing all carts
func ListCartsHandler(c *gin.Context) {
	var carts []models.Cart
	if err := db.DB.Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve carts"})
		return
	}

	c.JSON(http.StatusOK, carts)
}
