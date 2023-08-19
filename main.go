package main

import (
	"embed"
	"os"
	"os/exec"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// type App struct {
// 	runtime *wails.Runtime
// }

// func NewApp() *App {
// 	return &App{}
// }

// func (app *App) startup() {
// 	// Code to run on startup
// }

// Botched implementation of integrating Vale in Wales.

func (app *App) ReadFile(filename string) (string, error) {
	// Implement the ReadFile function to read the content of a file
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (app *App) RunVale(content string) (string, error) {
	// RunVale function to run the Vale command and capture the output
	cmd := exec.Command("vale", "--output", "line", "--", "-")
	cmd.Stdin = strings.NewReader(content)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "459svelte",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
