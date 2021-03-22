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

	// User sign-up and log-in API
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	// User post and get todo_list API
	routes.TodoListRoutes(router)

	port := "80"
	router.Run(":" + port)
}
