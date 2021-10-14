package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
//Create mongo connection
var MongoC = ConnectBD()
//Uri to connect to own BBDD
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@mongocb.c0efb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

// ConnectBD is the function to connect with the uri of mongoDB
func ConnectBD() *mongo.Client{
	client,err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connected to BBDD.")
	return client
}

//ConnectionCheck verify the status of bbdd with a "Ping"
func ConnectionCheck() int  {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 1
}