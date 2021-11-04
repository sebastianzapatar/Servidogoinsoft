package bd

import (
	"context"
	"fmt"
	"ingsoft/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func InsertarRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	bd := MongoCN.Database("baseingsoft")
	col := bd.Collection("usuarios")
	u.Password, _ = EncriptarPasword(u.Password)
	fmt.Println(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
func EncriptarPasword(password string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	return string(bytes), err
}
