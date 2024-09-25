package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/username/back_nkah_ecity/db"
	"github.com/username/back_nkah_ecity/models"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// GetGroups handles retrieving the list of groups
func GetGroups(c *gin.Context) {
	var groups []models.Group
	cursor, err := db.MongoClient.Database("DBEcity").Collection("groups").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve groups"})
		return
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var group models.Group
		if err := cursor.Decode(&group); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode group data"})
			return
		}
		groups = append(groups, group)
	}

	c.JSON(http.StatusOK, groups)
}

// CreateGroup handles creating a new group
func CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.MongoClient.Database("DBEcity").Collection("groups").InsertOne(context.TODO(), group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group created successfully"})
}
