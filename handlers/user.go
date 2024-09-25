package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/username/back_nkah_ecity/db"
	"github.com/username/back_nkah_ecity/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetUserProfile handles the retrieval of a user's profile by ID
func GetUserProfile(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = db.MongoClient.Database("DBEcity").Collection("users").FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserProfile handles the update of a user's profile
func UpdateUserProfile(c *gin.Context) {
	userID := c.Param("id")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": user}
	_, err = db.MongoClient.Database("DBEcity").Collection("users").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
