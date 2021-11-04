package routes

import (
	"encoding/json"
	"ingsoft/bd"
	"ingsoft/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	/*
		Se valida que no nos manden nada vacio
		Se valida que el email no exista
	*/
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "No se envio el campo email", 400)
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password no tiene la longitud indicada", 400)
		return
	}
	_, encontrado, _ := bd.CheckUser(t.Email)
	if encontrado == true {
		http.Error(w, "Usuario ya registrado", 400)
		return
	}
	_, status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Erro al insertar en la base datos"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Erro al insertar en la base datos", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
