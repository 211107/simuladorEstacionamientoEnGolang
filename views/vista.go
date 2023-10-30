package views

import (
	"simulador/models"
	"time" // Asegúrate de que esta línea esté presente para usar time

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

	// Cargar la imagen del carrito
	autoImagen := canvas.NewImageFromFile("assets/car.png")
	autoImagen.FillMode = canvas.ImageFillOriginal

	vista := container.NewHBox(layout.NewSpacer(), estacionamientoImagen, autoImagen, layout.NewSpacer())
	iniciarBoton := widget.NewButton("Iniciar", func() {
		// Iniciar la simulación de llegada continua de vehículos desde el paquete models
		go models.SimularEstacionamiento(100, 20)

		// Coordenadas iniciales del carro
		posX := float32(0)
		posY := float32(0)

		// Coordenadas de estacionamiento
		cajonX := float32(130) // Coordenada X del cajón
		cajonY := float32(70)  // Coordenada Y del cajón

		// Crear una función para mover el carro hacia el cajón
		moverHaciaCajon := func() {
			for {
				if posX < cajonX {
					posX += 2
				} else if posY < cajonY {
					posY += 2
				} else {
					// El carro ha llegado al cajón
					break
				}

				// Actualizar la posición de la imagen del carro en la vista
				autoImagen.Move(fyne.NewPos(posX, posY))
				myWindow.Canvas().Refresh(autoImagen)
				time.Sleep(100 * time.Millisecond) // Controla la velocidad de movimiento
			}
		}

		// Iniciar la función de movimiento hacia el cajón en segundo plano
		go moverHaciaCajon()
	})

	vistaConBoton := container.NewVBox(iniciarBoton, vista)

	myWindow.SetContent(vistaConBoton)
	myWindow.Resize(fyne.NewSize(500, estacionamientoImagen.Size().Height))
	myWindow.ShowAndRun()
}
