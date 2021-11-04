package bd

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://elcasique:TGXWeVbd1e97mO1q@cluster0.mwffh.mongodb.net/test")

func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	return client
}
func CheckConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	fmt.Println("Conecto")
	return 1
}
