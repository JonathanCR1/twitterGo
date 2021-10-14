package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//User Model of users in the BBDD
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"lastname" json:"lastname,omitempty"`
	BornDate  time.Time          `bson:"born" json:"born,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Bio       string             `bson:"bio" json:"bio,omitempty"`
	Ubication string             `bson:"ubication" json:"ubication,omitempty"`
	WebSite   string             `bson:"webSite" json:"webSite,omitempty"`
}
