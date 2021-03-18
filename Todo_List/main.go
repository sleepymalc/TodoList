package main

import (
	middleware "Todo_List/middleware"
	routes "Todo_List/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.Use(middleware.Authentication())

	// User sign-up and log-in API
	routes.UserRoutes(router)

	// User post and get todo_list API
	routes.TodoListRoutes(router)

	port := "8080"
	router.Run(":" + port)
}
