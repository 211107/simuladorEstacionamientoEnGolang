// models/auto.go
package models

import (
	"time"
	"math/rand"
)

type Auto struct {
	Id int
}

func CrearAuto(id int, estacionamiento *Estacionamiento, entrando chan<- Auto, saliendo chan<- Auto) {
	auto := Auto{Id: id} // Aquí está la corrección

	go func() {
		// Intentar entrar al estacionamiento
		entrando <- auto

		// Esperar un tiempo aleatorio entre 1 y 5 segundos
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

		// Intentar salir del estacionamiento
		saliendo <- auto
	}()
}