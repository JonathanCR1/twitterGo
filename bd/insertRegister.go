package bd

import (
	"context"
	"github.com/JonathanCR1/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//RegisterUser function to register a new user
func RegisterUser(user models.User)(string,bool,error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("twitter")
	col := db.Collection("users")
	user.Password,_ = EncriptPassw(user.Password)
	result, err := col.InsertOne(ctx,user)
	if err != nil{
		return "",false,err
	}
	ObjID,_ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(),true,nil
}

