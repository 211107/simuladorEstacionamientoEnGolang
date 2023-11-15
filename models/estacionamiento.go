package models

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Estacionamiento struct {
	espacios      chan int
	puerta        *sync.Mutex
	espaciosArray [20]bool
}

func NewEstacionamiento(espacios chan int, puertaMu *sync.Mutex) *Estacionamiento {
	return &Estacionamiento{
		espacios:      espacios,
		puerta:        puertaMu,
		espaciosArray: [20]bool{},
	}
}

func (p *Estacionamiento) GetEspacios() chan int {
	return p.espacios
}

func (p *Estacionamiento) GetPuertaMu() *sync.Mutex {
	return p.puerta
}

func (p *Estacionamiento) GetEspaciosArray() [20]bool {
	return p.espaciosArray
}

func (p *Estacionamiento) SetEspaciosArray(espaciosArray [20]bool) {
	p.espaciosArray = espaciosArray
}

func (p *Estacionamiento) ColaSalida(contenedor *fyne.Container, imagen *canvas.Image) {
    // Mueve la imagen al lugar deseado
    imagen.Move(fyne.NewPos(80, 20))

    // Agrega la imagen al contenedor y actualiza la interfaz
    contenedor.Add(imagen)
    contenedor.Refresh()

    // Espera un momento (ajusta el tiempo según sea necesario)
    time.AfterFunc(2*time.Second, func() {
        // Elimina la imagen del contenedor y actualiza la interfaz
        contenedor.Remove(imagen)
        contenedor.Refresh()
    })
}


