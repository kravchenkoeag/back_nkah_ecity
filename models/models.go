package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents the structure for a user in the system
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	Location  string             `bson:"location" json:"location"`
	Interests []string           `bson:"interests" json:"interests"`
}

// Group represents the structure for a group in the system
type Group struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string               `bson:"name" json:"name" validate:"required"`
	Description string               `bson:"description" json:"description"`
	Members     []primitive.ObjectID `bson:"members" json:"members"`
	Subgroups   []primitive.ObjectID `bson:"subgroups" json:"subgroups"` // New field for storing subgroups
}

// Announcement represents the structure for an announcement in the system
type Announcement struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Location    string             `bson:"location" json:"location"`
	AuthorID    primitive.ObjectID `bson:"author_id" json:"author_id"`
}

// Event represents the structure for an event in the system
type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date"`
	Location    string             `bson:"location" json:"location"`
	OrganizerID primitive.ObjectID `bson:"organizer_id" json:"organizer_id"`
}
