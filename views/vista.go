package views

import (
	"simulador/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func IniciarVentana() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Estacionamiento")
	myWindow.SetFixedSize(true)

	// Cargar la imagen del estacionamiento
	estacionamientoImagen := canvas.NewImageFromFile("assets/fondo.png")
	estacionamientoImagen.FillMode = canvas.ImageFillOriginal

	vista := container.NewHBox(layout.NewSpacer(), estacionamientoImagen, layout.NewSpacer())

	iniciarBoton := widget.NewButton("Iniciar", func() {
		// Iniciar la simulación de llegada continua de vehículos desde el paquete models
		cantidadCarros := 10
		espaciadoX := 44

		for i := 0; i < cantidadCarros; i++ {
			autoImagen := canvas.NewImageFromFile("assets/car.png")
			autoImagen.FillMode = canvas.ImageFillOriginal

			posX := float32(0)
			posY := float32(10)
			cajonX := float32(100 + espaciadoX*i)
			cajonY := float32(70)

			moverHaciaCajon := func() {
				for {
					if posX < cajonX {
						posX += 2
					} else if posY < cajonY {
						posY += 2
					} else {
						break
					}

					autoImagen.Move(fyne.NewPos(posX, posY))
					myWindow.Canvas().Refresh(autoImagen)
					time.Sleep(10 * time.Millisecond)
				}

				// Después de llegar al cajón, mover hacia abajo
				for {
					if posY < 60 { // Cambia esta condición según tu necesidad
						posY += 2
					} else {
						break
					}

					autoImagen.Move(fyne.NewPos(posX, posY))
					myWindow.Canvas().Refresh(autoImagen)
					time.Sleep(10 * time.Millisecond)
				}
			}

			vista.Add(autoImagen)
			go moverHaciaCajon()
		}

		go models.SimularEstacionamiento(cantidadCarros, cantidadCarros)
	})

	vistaConBoton := container.NewVBox(iniciarBoton, vista)

	myWindow.SetContent(vistaConBoton)
	myWindow.Resize(fyne.NewSize(500, estacionamientoImagen.Size().Height))
	myWindow.ShowAndRun()
}
