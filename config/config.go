package main

import (
	"backend/db"
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()
	defer db.Close()

	router := gin.Default()

	// User endpoints
	router.POST("/users", handlers.CreateUserHandler)
	router.GET("/users", handlers.ListUsersHandler)
	router.POST("/users/login", handlers.LoginUserHandler)

	// Item endpoints
	router.POST("/items", handlers.CreateItemHandler)
	router.GET("/items", handlers.ListItemsHandler)

	// Cart endpoints
	router.POST("/carts", handlers.CreateCartHandler)
	router.GET("/carts", handlers.ListCartsHandler)

	// Order endpoints
	router.POST("/orders", handlers.CreateOrderHandler)
	router.GET("/orders", handlers.ListOrdersHandler)

	router.Run(":8080")
}
