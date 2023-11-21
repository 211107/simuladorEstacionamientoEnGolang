package scenes

import (
	"simulador/models"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gonum.org/v1/gonum/stat/distuv"
)

type EscenaPrincipal struct {
	ventana fyne.Window
}

func NuevaEscenaPrincipal(ventana fyne.Window) *EscenaPrincipal {
	return &EscenaPrincipal{
		ventana: ventana,
	}
}


func (s *EscenaPrincipal) Mostrar() {
	fondoEstacionamiento := canvas.NewImageFromFile("assets/fon1.png")
	fondoEstacionamiento.Resize(fyne.NewSize(690, 400))
	fondoEstacionamiento.Move(fyne.NewPos(0, 0))

	contenedor := container.NewWithoutLayout()
	contenedor.Add(fondoEstacionamiento)

	boton := widget.NewButton("Iniciar", func() {
		go s.Ejecutar() // Inicia la simulación en una goroutine.
	})

	vbox := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		boton,
	)

	// Agrega el contenedor vertical al contenedor principal.
	contenedor.Add(vbox)
	s.ventana.SetContent(contenedor)
}

func (s *EscenaPrincipal) Ejecutar() {
	p := models.NuevoEstacionamiento(make(chan int, 20), &sync.Mutex{})
	contenedor := s.ventana.Content().(*fyne.Container)

	var wg sync.WaitGroup

	// Inicia 100 goroutines para simular vehículos ingresando al estacionamiento.
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			// Crea un nuevo vehículo con un ID único.
			vehiculo := models.NuevoVehiculo(id)
			imagen := vehiculo.ObtenerImagenEntrada()
			imagen.Resize(fyne.NewSize(30, 50))
			imagen.Move(fyne.NewPos(40, -10))
			contenedor.Add(imagen)
			contenedor.Refresh()
			vehiculo.Iniciar(p, contenedor, &wg)
			time.Sleep(time.Millisecond * 200)
		}(i)
		poisson := generarPoisson(2)
		time.Sleep(time.Second * time.Duration(poisson))
	}
	// Espera a que todas las goroutines
	wg.Wait()
}

// generarPoisson genera un número aleatorio según la distribución de Poisson con un parámetro lambda dado.
func generarPoisson(lambda float64) float64 {
	poisson := distuv.Poisson{Lambda: lambda, Src: nil}
	return poisson.Rand()
}
