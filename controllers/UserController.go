package controllers

import (
	"context"
	"crud_go/models"
	"crud_go/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []bson.M
	cursor, err := services.DB.Collection("users").Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user bson.M

		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

//postgres
//func GetUsers(c *gin.Context) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	query := "SELECT id, name, password FROM users"
//	rows, err := services.Pool.Query(ctx, query)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
//		return
//	}
//	defer rows.Close()
//
//	var users []models.Users
//	for rows.Next() {
//		var user models.Users
//		if err := rows.Scan(&user.ID, &user.Name, &user.Password); err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process user data"})
//			return
//		}
//		users = append(users, user)
//	}
//
//	c.JSON(http.StatusOK, users)
//}

func GetUserDetails(c *gin.Context) {
	c.JSON(200, gin.H{"success": "Detail users!" + c.Param("id")})
}

func AddUsers(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hasPassword)
	user.ID = bson.NewObjectID()

	_, err := services.DB.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func EditUsers(c *gin.Context) {

	var requestBody map[string]interface{}
	// Bind JSON ke dalam map
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	fmt.Println(objID)

	// Update data di MongoDB
	update := bson.M{
		"$set": bson.M{
			"name": requestBody["name"],
		},
	}

	_, err = services.DB.Collection("users").UpdateByID(context.TODO(), objID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUsers(c *gin.Context) {
	userID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	filter := bson.M{"_id": objID}
	_, err = services.DB.Collection("users").DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Deleted successfully"})
}
