package models

import (
	"time"

	"fmt"
	"sync"
    "math/rand"
    
)

type Estacionamiento struct {
	Cajones []int
	Mux     sync.Mutex
}

func (e *Estacionamiento) Entrar(auto Auto) {
	e.Mux.Lock()
	defer e.Mux.Unlock()

	for i := 0; i < len(e.Cajones); i++ {
		if e.Cajones[i] == 0 {
			e.Cajones[i] = auto.Id
			fmt.Printf("Auto %d ha entrado al cajón %d\n", auto.Id, i)
			break
		}
	}
}

func (e *Estacionamiento) Salir(auto Auto) {
	e.Mux.Lock()
	defer e.Mux.Unlock()

	for i := 0; i < len(e.Cajones); i++ {
		if e.Cajones[i] == auto.Id {
			e.Cajones[i] = 0
			fmt.Printf("Auto %d ha salido del cajón %d\n", auto.Id, i)
			break
		}
	}
}


func SimularEstacionamiento(numAutos int, numCajones int) {
	cajonesEstacionamiento := make([]int, numCajones)
	estacionamiento := &Estacionamiento{Cajones: cajonesEstacionamiento}

	entrando := make(chan Auto, numAutos)
	saliendo := make(chan Auto, numAutos)

	go func() {
		for {
			select {
			case auto := <-entrando:
				estacionamiento.Entrar(auto)
			case auto := <-saliendo:
				estacionamiento.Salir(auto)
			}
		}
	}()

	for i := 1; i <= numAutos; i++ {
		go CrearAuto(i, estacionamiento, entrando, saliendo)
		
        // Simulación de llegadas a través de distribución poison.
        lambda := 0.2
        poisson := rand.ExpFloat64() / lambda
        time.Sleep(time.Duration(poisson) * time.Second)
	}

	time.Sleep(10 * time.Second)
}