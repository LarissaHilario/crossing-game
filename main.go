// main.go
package main

import (
	"juego-pollo/scenes"

	"fyne.io/fyne/v2/app"
)

const (
    gameWidth     = 300
    gameHeight    = 500
    playerSize    = 50
    obstacleSize  = 70
)

func main() {
    myapp:= app.New()
    mywindow:= myapp.NewWindow()
    
    menuScene := scenes.NewMenuScene()
    menuScene.Start()
}
