package main

import (
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// Встраиваем приватный ключ
//
//go:embed id_rsa
var privateKey []byte

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла:", err)
	}

	// Проверяем наличие необходимых переменных окружения
	if err := checkEnvVars(); err != nil {
		log.Fatal(err)
	}

	// Create an instance of the app structure
	app := NewApp(privateKey)

	// Create application with options
	err := wails.Run(&options.App{
		Title:    "SHH-клиент Minecraft",
		MaxWidth: 450,
		MinWidth: 450,
		Width:    450,

		MaxHeight: 600,
		MinHeight: 600,
		Height:    600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

// Проверяем наличие обязательных переменных окружения
func checkEnvVars() error {
	requiredVars := []string{"SERVER_USER", "SERVER_HOST", "SERVER_PORT", "LOCAL_PORT"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("переменная окружения %s не найдена", v)
		}
	}
	return nil
}
