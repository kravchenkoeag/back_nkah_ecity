package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kravchenkoeag/back_nkah_ecity/db"
	"github.com/kravchenkoeag/back_nkah_ecity/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Check if a group with the same name already exists
	var existingGroup models.Group
	err := db.MongoClient.Database("DBEcity").Collection("groups").FindOne(context.TODO(), bson.M{"name": group.Name}).Decode(&existingGroup)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group with this name already exists"})
		return
	}

	// Insert the new group into the MongoDB collection
	_, err = db.MongoClient.Database("DBEcity").Collection("groups").InsertOne(context.TODO(), group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group created successfully"})
}

// UpdateGroup handles renaming or updating the description of a group
func UpdateGroup(c *gin.Context) {
	groupID := c.Param("id")
	var updateData struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectId, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	update := bson.M{}
	if updateData.Name != "" {
		// Check if a group with the same name already exists
		var existingGroup models.Group
		err := db.MongoClient.Database("DBEcity").Collection("groups").FindOne(context.TODO(), bson.M{"name": updateData.Name}).Decode(&existingGroup)
		if err == nil && existingGroup.ID != objectId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Group with this name already exists"})
			return
		}
		update["name"] = updateData.Name
	}
	if updateData.Description != "" {
		update["description"] = updateData.Description
	}

	_, err = db.MongoClient.Database("DBEcity").Collection("groups").UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.M{"$set": update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group updated successfully"})
}

// DeleteGroup handles deleting a group if it has no subgroups or members
func DeleteGroup(c *gin.Context) {
	groupID := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID"})
		return
	}

	// Check if the group has any subgroups
	var subGroup models.Group
	err = db.MongoClient.Database("DBEcity").Collection("groups").FindOne(context.TODO(), bson.M{"subgroups": objectId}).Decode(&subGroup)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete group with existing subgroups"})
		return
	}

	// Check if the group has any members
	var group models.Group
	err = db.MongoClient.Database("DBEcity").Collection("groups").FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&group)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	if len(group.Members) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete group with registered members"})
		return
	}

	// Delete the group if there are no subgroups or members
	_, err = db.MongoClient.Database("DBEcity").Collection("groups").DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group deleted successfully"})
}
