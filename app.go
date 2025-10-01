package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"next-drive/constants"
	"os"

	"github.com/Zorgatone/next-drive-api/client"
	"github.com/wailsapp/wails/v2/pkg/logger"
)

// App struct
type App struct {
	ctx       context.Context
	appLogger logger.Logger
	apiClient *client.ApiClient
}

// NewApp creates a new App application struct
func NewApp(appLogger logger.Logger) *App {
	return &App{
		appLogger: appLogger,
		apiClient: nil,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) (err error) {
	a.ctx = ctx

	host := os.Getenv("NC_HOST")

	if host == "" {
		return errors.New("missing env NC_HOST")
	}

	apiClient, err := client.NewApiClient(&url.URL{
		Scheme: "https",
		Host:   host,
	}, constants.UserAgentHeaderValue)

	if err != nil {
		return err
	}

	a.apiClient = apiClient

	username := os.Getenv("NC_USERNAME")

	if username == "" {
		return errors.New("missing env NC_USERNAME")
	}

	password := os.Getenv("NC_PASSWORD")

	if password == "" {
		return errors.New("missing env NC_PASSWORD")
	}

	err = a.Login(username, password)

	if err != nil {
		return err
	}

	a.appLogger.Info("logged in")

	return nil
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(username string, password string) (err error) {
	err = a.apiClient.Login(client.LoginCredentials{
		User:     username,
		Password: password,
	})

	if err != nil {
		return err
	}

	return nil
}
