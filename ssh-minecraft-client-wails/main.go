package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

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

// Встраиваем .env файл
//
//go:embed .env
var envFile []byte

func main() {
	// Загружаем переменные окружения из встроенного .env файла
	loadEmbeddedEnv()

	// Проверяем наличие необходимых переменных окружения
	if err := checkEnvVars(); err != nil {
		log.Fatal(err)
	}

	// Create an instance of the app structure
	app := NewApp(privateKey)

	// Create application with options
	err := wails.Run(&options.App{
		Title:    "SSH-клиент Minecraft",
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

// Загрузка переменных окружения из встроенного .env файла
func loadEmbeddedEnv() {
	envData := string(envFile)
	lines := strings.Split(envData, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue // Пропускаем пустые строки и комментарии
		}

		// Разделяем строку на ключ и значение
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Пропускаем строки без корректного формата
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Устанавливаем переменную окружения
		err := os.Setenv(key, value)
		if err != nil {
			log.Fatalf("Ошибка установки переменной окружения: %s", err)
		}
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
