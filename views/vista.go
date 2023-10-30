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

	// Cargar la imagen del primer carro
	autoImagen := canvas.NewImageFromFile("assets/car.png")
	autoImagen.FillMode = canvas.ImageFillOriginal

	// Cargar la imagen del segundo carro
	autoImagen2 := canvas.NewImageFromFile("assets/car.png") // Cambiar la imagen si es diferente
	autoImagen2.FillMode = canvas.ImageFillOriginal

	vista := container.NewHBox(layout.NewSpacer(), estacionamientoImagen, autoImagen, autoImagen2, layout.NewSpacer())
	iniciarBoton := widget.NewButton("Iniciar", func() {
		// Iniciar la simulación de llegada continua de vehículos desde el paquete models
		go models.SimularEstacionamiento(2, 20) // Cambiar a 2 carros

		// Coordenadas iniciales y cajón del primer carro
		posX := float32(0)
		posY := float32(0)
		cajonX := float32(107)
		cajonY := float32(60)

		// Coordenadas iniciales y cajón del segundo carro
		posX2 := float32(0)
		posY2 := float32(0)
		cajonX2 := float32(138) // Cambiar a la posición X del segundo cajón
		cajonY2 := float32(60)  // Cambiar a la posición Y del segundo cajón

		// Crear una función para mover el primer carro hacia el cajón
		moverHaciaCajon := func() {
			for {
				if posX < cajonX {
					posX += 2
				} else if posY < cajonY {
					posY += 2
				} else {
					// El primer carro ha llegado al primer cajón
					break
				}

				// Actualizar la posición de la imagen del primer carro en la vista
				autoImagen.Move(fyne.NewPos(posX, posY))
				myWindow.Canvas().Refresh(autoImagen)
				time.Sleep(100 * time.Millisecond)
			}
		}

		// Crear una función para mover el segundo carro hacia el cajón
		moverHaciaCajon2 := func() {
			for {
				if posX2 < cajonX2 {
					posX2 += 2
				} else if posY2 < cajonY2 {
					posY2 += 2
				} else {
					// El segundo carro ha llegado al segundo cajón
					break
				}

				// Actualizar la posición de la imagen del segundo carro en la vista
				autoImagen2.Move(fyne.NewPos(posX2, posY2))
				myWindow.Canvas().Refresh(autoImagen2)
				time.Sleep(100 * time.Millisecond)
			}
		}

		// Iniciar la función de movimiento hacia el cajón para ambos carros en segundo plano
		go moverHaciaCajon()
		go moverHaciaCajon2() // Agregar el segundo carro
	})

	vistaConBoton := container.NewVBox(iniciarBoton, vista)

	myWindow.SetContent(vistaConBoton)
	myWindow.Resize(fyne.NewSize(500, estacionamientoImagen.Size().Height))
	myWindow.ShowAndRun()
}
