package routes

import (
	controller "Todo_List/controllers"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())
}

//TodoListRoutes function
func TodoListRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/todo_list", controller.TodoListPost())
	incomingRoutes.GET("/users/todo_list", controller.TodoListGet())
	incomingRoutes.DELETE("/users/todo_list", controller.TodoListDelete())
}
