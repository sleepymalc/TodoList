package controllers

import (
	"context"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"Todo_List/database"

	helper "Todo_List/helpers"
	"Todo_List/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "userInfo")

var validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("password is incorrect")
		check = false
	}

	return check, msg
}

//CreateUser is the api used to get a single user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var todo_list models.TodoList

		if userid, exist := c.GetPostForm("user_id"); exist && userid != "" {
			user.User_id = userid
			todo_list.User_id = userid
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username can't be empty"})
			return
		}

		if password, exist := c.GetPostForm("password"); exist && password != "" {
			user.Password = &password
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password can't be empty"})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"user_id": user.User_id})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the user"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "this user id already exists"})
			return
		}

		//user info
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		token, refreshToken, _ := helper.GenerateAllTokens(user.User_id)
		user.Token = &token
		user.Refresh_token = &refreshToken

		//todolist userInfo
		todo_list.ID = primitive.NewObjectID()
		todo_list.Todo_list = make([]string, 0, 10)

		resultInsertionNumberUserInfo, insertErr := userCollection.InsertOne(ctx, user)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		_, insertErr = todoListCollection.InsertOne(ctx, todo_list)
		if insertErr != nil {
			msg := fmt.Sprintf("Todo list item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumberUserInfo)
	}
}

//Login is the api used to get a single user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		if userid, exist := c.GetPostForm("user_id"); exist && userid != "" {
			user.User_id = userid
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username can't be empty"})
			return
		}

		if password, exist := c.GetPostForm("password"); exist && password != "" {
			user.Password = &password
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Password can't be empty"})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"user_id": user.User_id}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
			return
		}

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
		defer cancel()
		if !passwordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, _ := helper.GenerateAllTokens(foundUser.User_id)

		helper.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, foundUser.Token)

	}
}
