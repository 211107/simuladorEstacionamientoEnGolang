package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Auto struct {
	id              int
	tiempoLim       time.Duration
	espacioAsignado int
	imagenEntrada   *canvas.Image
	imagenEspera    *canvas.Image
	imagenSalida    *canvas.Image
}

func NewAuto(id int) *Auto {
	imagenEntrada := canvas.NewImageFromURI(storage.NewFileURI("./assets/carEntra.png"))
	imagenSalida := canvas.NewImageFromURI(storage.NewFileURI("./assets/carSalida.png"))
	return &Auto{
		id:              id,
		tiempoLim:       time.Duration(rand.Intn(50)+50) * time.Second,
		espacioAsignado: 0,
		imagenEntrada:   imagenEntrada,
		imagenSalida:    imagenSalida,
	}
}

func (a *Auto) Entrar(p *Estacionamiento, contenedor *fyne.Container) {
	p.GetEspacios() <- a.GetId()
	p.GetPuertaMu().Lock()

	espacios := p.GetEspaciosArray()
	const (
		columnasPorGrupo  = 10
		espacioHorizontal = 57
		espacioVertical   = 320
	)

	for i := 0; i < len(espacios); i++ {
		if !espacios[i] {
			espacios[i] = true
			a.espacioAsignado = i

			fila := i / (columnasPorGrupo * 1)
			columna := i % (columnasPorGrupo * 1)

			if columna >= columnasPorGrupo {
				columna += 1
			}

			x := float32(133 + columna*espacioHorizontal)
			if columna >= columnasPorGrupo {

			}
			y := float32(10 + fila*espacioVertical)

			a.imagenEntrada.Move(fyne.NewPos(x, y))
			break
		}
	}

	p.SetEspaciosArray(espacios)

	p.GetPuertaMu().Unlock()
	contenedor.Refresh()
	fmt.Printf("Auto %d ocupó el lugar %d.\n", a.GetId(), a.espacioAsignado)
	 // Pausa de 5 segundos antes de continuar.
	 time.Sleep(5 * time.Second)
}

func (a *Auto) Salir(p *Estacionamiento, contenedor *fyne.Container) {
    // Espera a que haya un espacio disponible en el estacionamiento.
    <-p.GetEspacios()
    // Bloquea la puerta para evitar que otros autos entren o salgan al mismo tiempo.
    p.GetPuertaMu().Lock()

    // Obtiene el array de espacios actual del estacionamiento.
    spacesArray := p.GetEspaciosArray()

    // Marca el espacio asignado por el auto como disponible.
    spacesArray[a.espacioAsignado] = false
	fmt.Printf("Auto %d salió. Espacio %d marcado como disponible.\n", a.GetId(), a.espacioAsignado)

    // Actualiza el array de espacios del estacionamiento.
    p.SetEspaciosArray(spacesArray)

    // Desbloquea la puerta para permitir que otros autos entren o salgan.
    p.GetPuertaMu().Unlock()

    // Elimina la imagen de espera del auto del contenedor.
    contenedor.Remove(a.imagenEntrada)
	contenedor.Refresh()

    // Ajusta el tamaño y la posición de la imagen de salida.
    a.imagenSalida.Resize(fyne.NewSize(30, 50))
    a.imagenSalida.Move(fyne.NewPos(90, 290))

    // Agrega la imagen de salida al contenedor y actualiza la interfaz.
    contenedor.Add(a.imagenSalida)
    contenedor.Refresh()

    // Realiza una animación de movimiento hacia arriba durante 10 iteraciones.
    for i := 0; i < 10; i++ {
        // Mueve la imagen hacia arriba.
        a.imagenSalida.Move(fyne.NewPos(a.imagenSalida.Position().X, a.imagenSalida.Position().Y-40))
		time.Sleep(time.Millisecond * 100)
    }

    // Elimina la imagen de salida del contenedor y actualiza la interfaz.
    contenedor.Remove(a.imagenSalida)
    contenedor.Refresh()  // Asegúrate de actualizar el contenedor

}


func (a *Auto) Iniciar(p *Estacionamiento, contenedor *fyne.Container, wg *sync.WaitGroup) {
	a.Avanzar(10)

	a.Entrar(p, contenedor)
	 // Pausa adicional de 5 segundos antes de iniciar el proceso de salida.
	 time.Sleep(5 * time.Second)
	  // Inicia un temporizador de 5 segundos antes de comenzar la salida.
	  timer := time.NewTimer(5 * time.Second)
	

	    // Usa un select para esperar tanto el temporizador como la finalización del temporizador.
		select {
		case <-timer.C:
			// El temporizador ha expirado, procede con la salida del auto.
			contenedor.Remove(a.imagenEntrada)
			contenedor.Refresh()
			contenedor.Add(a.imagenSalida)
			contenedor.Refresh()
			//a.imagenSalida.Resize(fyne.NewSize(50, 30))
			p.ColaSalida(contenedor, a.imagenEntrada)
			a.Salir(p, contenedor)
			
			wg.Done()
		}
	}

func (a *Auto) Avanzar(pasos int) {
	for i := 0; i < pasos; i++ {
		a.imagenEntrada.Move(fyne.NewPos(a.imagenEntrada.Position().X, a.imagenEntrada.Position().Y+20))
		time.Sleep(time.Millisecond * 100)
	}
}

func (a *Auto) GetId() int {
	return a.id
}

// Agrega estos métodos para proporcionar acceso a los campos necesarios
func (a *Auto) GetTiempoLim() time.Duration {
    return a.tiempoLim
}

func (a *Auto) GetImagenEntrada() *canvas.Image {
    return a.imagenEntrada
}

func (a *Auto) GetImagenSalida() *canvas.Image {
    return a.imagenSalida
}

func (a *Auto) GetImagenEspera() *canvas.Image {
    return a.imagenEspera
}



