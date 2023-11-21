package models

import (
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)


type Estacionamiento struct {
	espacios      chan int      // Canal para representar los espacios disponibles
	puerta        *sync.Mutex  
	arrayEspacios [20]bool      
}

// NuevoEstacionamiento crea una nueva instancia de Estacionamiento.
func NuevoEstacionamiento(espacios chan int, puertaMu *sync.Mutex) *Estacionamiento {
	return &Estacionamiento{
		espacios: espacios,
		puerta:   puertaMu,
	}
}

// ObtenerEspacio devuelve el canal de espacios del estacionamiento.
func (e *Estacionamiento) ObtenerEspacio() chan int {
	return e.espacios
}

// ObtenerPuerta devuelve el puntero al Mutex utilizado para gestionar el acceso a recursos compartidos.
func (e *Estacionamiento) ObtenerPuerta() *sync.Mutex {
	return e.puerta
}

// ObtenerArrayEspacios devuelve el array que representa el estado de ocupación de los espacios.
func (e *Estacionamiento) ObtenerArrayEspacios() [20]bool {
	return e.arrayEspacios
}

// EstablecerArrayEspacios establece el array que representa el estado de ocupación de los espacios.
func (e *Estacionamiento) EstablecerArrayEspacios(arrayEspacios [20]bool) {
	e.arrayEspacios = arrayEspacios
}


func (e *Estacionamiento) ColaSalida(contenedor *fyne.Container, imagen *canvas.Image) {
	imagen.Move(fyne.NewPos(80, 20))
	contenedor.Add(imagen)
	contenedor.Refresh()

	// Utiliza un temporizador para quitar la imagen
	time.AfterFunc(2*time.Second, func() {
		contenedor.Remove(imagen)
		contenedor.Refresh()
	})
}
