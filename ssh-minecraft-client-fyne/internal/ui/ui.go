package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"minecraft-client/internal/ssh"
)

var MinecraftLogo []byte
var MinecraftIcon []byte
var errorLabel *widget.Label

// RunUI Создание и запуск интерфейса
func RunUI(privateKey []byte) {
	myApp := app.New()

	// Устанавливаем иконку приложения Minecraft
	iconResource := fyne.NewStaticResource("minecraft_icon", MinecraftIcon)
	myApp.SetIcon(iconResource)

	myWindow := myApp.NewWindow("Подключение к серверу Minecraft")

	// Логотип Minecraft (Пиксельное яблоко)
	logoImage := canvas.NewImageFromResource(fyne.NewStaticResource("minecraft_logo", MinecraftLogo))
	logoImage.FillMode = canvas.ImageFillContain
	logoImage.SetMinSize(fyne.NewSize(150, 150))

	// Создаем текстовые элементы
	errorLabel = widget.NewLabel("")
	serverLabel := widget.NewLabelWithStyle("Для подключения используйте адрес сервера:", fyne.TextAlignCenter, fyne.TextStyle{})
	serverHostLabel := canvas.NewText("localhost", color.White)
	serverHostLabel.Alignment = fyne.TextAlignCenter
	serverHostLabel.TextSize = 24

	infoLabel := canvas.NewText("При подключении будет создан ssh-туннель с сервером Minecraft на локальный порт 25565", color.Gray{Y: 100})
	infoLabel.Alignment = fyne.TextAlignCenter
	infoLabel.TextSize = 10

	// Создаем кастомную кнопку с динамическим фоном
	var toggleButton *widget.Button

	// Создаем прямоугольник для фона кнопки
	buttonBackground := canvas.NewRectangle(&color.RGBA{4, 184, 108, 255}) // 04B86C - зеленый цвет для "ПОДКЛЮЧИТЬСЯ"

	// Статусы подключения и ошибки
	statusStable := canvas.NewText("СОЕДИНЕНИЕ СТАБИЛЬНО", color.RGBA{4, 184, 108, 255})
	statusStable.Alignment = fyne.TextAlignCenter
	statusStable.Hide()

	statusError := canvas.NewText("ОШИБКА", color.RGBA{197, 56, 42, 255})
	statusError.Alignment = fyne.TextAlignCenter
	statusError.Hide()

	toggleButton = widget.NewButton("ПОДКЛЮЧИТЬСЯ", func() {
		if !ssh.IsConnected {
			if err := ssh.ConnectSSH(privateKey); err != nil {
				statusStable.Hide() // Скрыть статус "СОЕДИНЕНИЕ СТАБИЛЬНО" при ошибке
				statusError.Show()  // Показать статус "ОШИБКА"
				displayError(err)
				return
			}
			statusStable.Show() // Показать статус "СОЕДИНЕНИЕ СТАБИЛЬНО"
			statusError.Hide()  // Скрыть статус "ОШИБКА"
			toggleButton.SetText("ОТКЛЮЧИТЬСЯ")
			buttonBackground.FillColor = color.RGBA{197, 56, 42, 255} // Красный фон
			canvas.Refresh(buttonBackground)
		} else {
			if err := ssh.DisconnectSSH(); err != nil {
				statusStable.Hide()
				statusError.Show()
				displayError(err)
				return
			}
			statusStable.Hide()
			statusError.Hide()
			toggleButton.SetText("ПОДКЛЮЧИТЬСЯ")
			buttonBackground.FillColor = color.RGBA{4, 184, 108, 255} // Зеленый фон
			canvas.Refresh(buttonBackground)
		}
	})

	// Упаковываем кнопку и фон в контейнер
	//buttonContainer := container.NewStack(buttonBackground, toggleButton)
	buttonContainer := container.NewStack(toggleButton)

	// Основной макет с логотипом, текстами и кнопкой
	content := container.NewVBox(
		logoImage,
		serverLabel,
		serverHostLabel,
		buttonContainer,
		infoLabel,
		statusStable,
		statusError,
	)

	content = container.NewPadded(container.NewPadded(content))

	// Создаем фоновый цвет темно-серого цвета
	background := canvas.NewRectangle(color.RGBA{40, 40, 40, 255})

	// Помещаем фон и содержимое в один контейнер
	myWindow.SetContent(container.NewMax(background, content))
	myWindow.Resize(fyne.NewSize(400, 500))
	myWindow.ShowAndRun()

	cleanup()
}

func displayError(err error) {
	errorLabel.SetText(fmt.Sprintf("Ошибка: %v", err))
	log.Println("Ошибка:", err)
}

func cleanup() {
	if ssh.IsConnected {
		if err := ssh.DisconnectSSH(); err != nil {
			log.Println("Ошибка при отключении SSH:", err)
		}
	}
}
