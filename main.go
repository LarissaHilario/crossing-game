// main.go
package main

import (
    "juego-pollo/scenes"  
)

const (
    gameWidth     = 300
    gameHeight    = 500
    playerSize    = 50
    obstacleSize  = 70
)

func main() {
    gameScene := scenes.NewGameScene()
    gameScene.Start()
}
