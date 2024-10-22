package ssh

import (
	"fmt"
	"os"
	"os/exec"
)

// Переменная для отслеживания состояния соединения
var sshCmd *exec.Cmd
var IsConnected bool = false

// ConnectSSH Подключение к SSH
func ConnectSSH(privateKey []byte) error {
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

	sshCmd = exec.Command("ssh", "-i", keyFile.Name(), "-N", "-L", fmt.Sprintf("%s:localhost:%s", localPort, localPort), fmt.Sprintf("%s@%s", serverUser, serverHost), "-p", serverPort)

	if err := sshCmd.Start(); err != nil {
		return fmt.Errorf("не удалось запустить команду SSH: %v", err)
	}

	go func() {
		sshCmd.Wait()
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
	}
	return nil
}
