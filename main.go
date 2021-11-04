package main

import (
	"ingsoft/bd"
	"ingsoft/handlers"
	"log"
)

func main() {
	if bd.CheckConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
