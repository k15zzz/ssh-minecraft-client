package main

import (
	_ "embed"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"minecraft-client/internal/ui"
	"os"
)

// Встраиваем логотип Minecraft и иконку
//
//go:embed assets/minecraft_logo.png
var minecraftLogo []byte

//go:embed assets/minecraft_logo.png
var minecraftIcon []byte

// Встраиваем приватный ключ
//
//go:embed assets/id_rsa
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

	// Передаем данные логотипа и иконки в интерфейс
	ui.MinecraftLogo = minecraftLogo
	ui.MinecraftIcon = minecraftIcon

	// Запускаем UI
	ui.RunUI(privateKey)
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
