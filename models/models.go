package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	Location  string             `bson:"location" json:"location"`
	Interests []string           `bson:"interests" json:"interests"`
}

type Group struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string               `bson:"name" json:"name" validate:"required"`
	Description string               `bson:"description" json:"description"`
	Members     []primitive.ObjectID `bson:"members" json:"members"`
}

type Announcement struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Location    string             `bson:"location" json:"location"`
	AuthorID    primitive.ObjectID `bson:"author_id" json:"author_id"`
}

type Event struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date"`
	Location    string             `bson:"location" json:"location"`
	OrganizerID primitive.ObjectID `bson:"organizer_id" json:"organizer_id"`
}
