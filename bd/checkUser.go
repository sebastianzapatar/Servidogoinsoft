package bd

import (
	"context"
	"fmt"
	"ingsoft/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(email string) (models.Usuario, bool, string) {
	fmt.Println(email, "CheckUser")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15) //se sale despues de 15 segundos
	defer cancel()
	bd := MongoCN.Database("baseingsoft")
	col := bd.Collection("usuarios")
	condicion := bson.M{"email": email}
	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.Id.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
