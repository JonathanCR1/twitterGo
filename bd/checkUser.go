package bd

import (
	"context"
	"github.com/JonathanCR1/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)
//UserExist This function return a bool for identify if the user-mail is resgistered
func UserExist(email string) (models.User ,bool, string){
	ctx,  cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()
	db := MongoC.Database("twitter")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var resultado models.User
	err := col.FindOne(ctx,condition).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil{
		return resultado,false,ID
	}
	return resultado, true, ID
}
