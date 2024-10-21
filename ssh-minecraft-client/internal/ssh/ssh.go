package ssh

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"os/exec"
	"syscall"
)

// Переменная для отслеживания состояния соединения
var sshCmd *exec.Cmd
var IsConnected bool = false

// ConnectSSH Подключение к SSH с таймаутом
func ConnectSSH(privateKey []byte, actx context.Context) error {
	serverUser := os.Getenv("SERVER_USER")
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	localPort := os.Getenv("LOCAL_PORT")

	keyFile, err := os.CreateTemp("", "id_rsa")
	if err != nil {
		return fmt.Errorf("не удалось создать временный файл для приватного ключа: %v", err)
	}

	if _, err := keyFile.Write(privateKey); err != nil {
		return fmt.Errorf("не удалось записать приватный ключ во временный файл: %v", err)
	}
	keyFile.Chmod(0600)
	keyFile.Close()

	sshCmd = exec.Command("ssh", "-i", keyFile.Name(), "-N", "-L",
		fmt.Sprintf("%s:localhost:%s", localPort, localPort),
		fmt.Sprintf("%s@%s", serverUser, serverHost), "-p", serverPort,
		"-o", "BatchMode=yes", "-o", "ExitOnForwardFailure=yes")

	if err := sshCmd.Start(); err != nil {
		return fmt.Errorf("не удалось запустить команду SSH: %v", err)
	}

	// Отслеживаем завершение команды и обрабатываем обрывы соединений
	go func() {
		err := sshCmd.Wait()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok && exitErr.Sys().(syscall.WaitStatus).Signal() == syscall.SIGKILL {
				runtime.EventsEmit(actx, "connectionEnd")
			} else {
				runtime.EventsEmit(actx, "connectionEndError", fmt.Sprintf("SSH соединение завершено с ошибкой: %v", err))
			}
		} else {
			runtime.EventsEmit(actx, "connectionEnd")
		}
		IsConnected = false
		os.Remove(keyFile.Name())
	}()

	IsConnected = true
	return nil
}

// DisconnectSSH Отключение SSH
func DisconnectSSH() error {
	if sshCmd != nil && IsConnected {
		if err := sshCmd.Process.Kill(); err != nil {
			return fmt.Errorf("ошибка отключения SSH: %v", err)
		}
		IsConnected = false
		fmt.Println("SSH соединение отключено.")
	}
	return nil
}
