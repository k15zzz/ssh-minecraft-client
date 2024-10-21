package main

import (
	"context"
	"fmt"
	"ssh-minecraft-client/internal/ssh"
)

type Request struct {
	Status bool   `json:"status"`
	Meta   string `json:"meta"`
}

// App struct
type App struct {
	ctx        context.Context
	privateKey []byte
}

// NewApp creates a new App application struct
func NewApp(privateKey []byte) *App {
	return &App{privateKey: privateKey}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	err := ssh.DisconnectSSH()
	if err != nil {
		fmt.Println(fmt.Sprintf("SSH соединение не завершено: %v", err))
	}
}

func (a *App) ConnectSSH() Request {
	err := ssh.ConnectSSH(a.privateKey, a.ctx)
	if err != nil {
		return Request{false, fmt.Sprintf("Ошибка подключения: %v", err)}
	}
	return Request{true, "Соединение установлено"}
}

func (a *App) DisconnectSSH() Request {
	err := ssh.DisconnectSSH()
	if err != nil {
		return Request{false, fmt.Sprintf("Ошибка отключения: %v", err)}
	}
	return Request{true, "Соединение разорвано"}
}

func (a *App) StatusConnection() bool {
	return ssh.IsConnected
}
