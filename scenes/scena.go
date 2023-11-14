// mainScene.go
package scenes

import (
	"simulador/models"
	"sync"
	"time"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gonum.org/v1/gonum/stat/distuv"
	
)

type MainScene struct {
	window fyne.Window
}

func NewMainScene(window fyne.Window) *MainScene {
	return &MainScene{
		window: window,
	}
}

func (s *MainScene) Show() {
    fondoEstacionamiento := canvas.NewImageFromFile("assets/fon1.png")
    fondoEstacionamiento.Resize(fyne.NewSize(690, 400))
    fondoEstacionamiento.Move(fyne.NewPos(0, 0))

    contenedor := container.NewWithoutLayout()
    contenedor.Add(fondoEstacionamiento)

    button := widget.NewButton("Iniciar", func() {
        go s.Run()
    })

    // Crear un contenedor vertical para posicionar el botón hacia abajo
    vbox := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
        layout.NewSpacer(), // Agrega espacio para mover el botón hacia la parte superior
        layout.NewSpacer(), // Agrega espacio para mover el botón hacia la parte superior
        layout.NewSpacer(), // Agrega espacio para mover el botón hacia la parte superior
        layout.NewSpacer(), // Agrega espacio para mover el botón hacia la parte superior
        button,
    )

    contenedor.Add(vbox) // Agrega el contenedor vertical al contenedor principal

    s.window.SetContent(contenedor)
}





func (s *MainScene) Run() {
	p := models.NewEstacionamiento(make(chan int, 20), &sync.Mutex{})
	contenedor := s.window.Content().(*fyne.Container) // Obtén el contenedor actual

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			auto := models.NewAuto(id)
			imagen := auto.GetImagenEntrada()
			imagen.Resize(fyne.NewSize(30, 50))
			imagen.Move(fyne.NewPos(40, -10))

			contenedor.Add(imagen)
			contenedor.Refresh()

			auto.Iniciar(p, contenedor, &wg)
		}(i)
		var poisson = generarPoisson(float64(2))
		time.Sleep(time.Second * time.Duration(poisson))
	}

	wg.Wait()
}

func generarPoisson(lambda float64) float64 {
	poisson := distuv.Poisson{Lambda: lambda, Src: nil}
	return poisson.Rand()
}
