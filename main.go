// main.go
package main

import (
	"juego-pollo/scenes"
)

const (
    gameWidth     = 500
    gameHeight    = 700
    playerSize    = 50
    obstacleSize  = 70
    
)

func main() {
	gameScene := scenes.NewGameScene()
    gameScene.Start()

}