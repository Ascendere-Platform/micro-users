package main

import (
	"log"
	"github.com/ascendere/micro-users/bd"
	"github.com/ascendere/micro-users/handlers"

)

func main (){
	if bd.ChequeoConnection() == 0{
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}