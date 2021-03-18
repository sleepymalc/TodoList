package controllers

import (
	"context"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"Todo_List/database"

	helper "Todo_List/helpers"
	"Todo_List/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoListCollection *mongo.Collection = database.OpenCollection(database.Client, "todoList")

//TodoListPost is the api used to add a todo_list element
func TodoListPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var insert_todo_list, user_todo_list models.TodoList

		var user_id string = c.GetString("user_id")

		if todo_list_array, exist := c.GetPostFormArray("todo_list"); exist && len(todo_list_array) != 0 {
			insert_todo_list.Todo_list = todo_list_array
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "todo list is empty"})
			return
		}

		err := todoListCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&user_todo_list)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "default todo list id not found"})
			return
		}

		user_todo_list.Todo_list = append(user_todo_list.Todo_list, insert_todo_list.Todo_list...)

		resultInsertionNumber, insertErr := todoListCollection.ReplaceOne(ctx, bson.M{"user_id": user_id}, user_todo_list)
		if insertErr != nil {
			msg := fmt.Sprintf("Todo list not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		token, refreshToken, _ := helper.GenerateAllTokens(user_id)

		helper.UpdateAllTokens(token, refreshToken, user_id)

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

//TodoListGet is the api used to get a todo_list
func TodoListGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var todo_list models.TodoList
		var user_id string = c.GetString("user_id")

		err := todoListCollection.FindOne(ctx, bson.M{"user_id": user_id}).Decode(&todo_list)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "todo list id not found"})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(user_id)

		helper.UpdateAllTokens(token, refreshToken, user_id)

		c.JSON(http.StatusOK, todo_list.Todo_list)
	}
}

//TodoListDelete is the api used to delete a todo_list element
func TodoListDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user_id string = c.GetString("user_id")
		var delete_element string

		if element, exist := c.GetPostForm("delete_element"); exist && element != "" {
			delete_element = element
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "todo element to be deleted is empty"})
			return
		}

		find_result := todoListCollection.FindOne(ctx, bson.M{"user_id": user_id})
		defer cancel()
		if find_result.Err() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user's todo list id not found"})
			return
		}

		resultInsertionNumber, insertErr := todoListCollection.UpdateOne(ctx, bson.M{"user_id": user_id}, bson.M{"$pull": bson.M{"todo_list": delete_element}})
		if insertErr != nil {
			msg := fmt.Sprintf("Todo list not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(user_id)

		helper.UpdateAllTokens(token, refreshToken, user_id)

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}
