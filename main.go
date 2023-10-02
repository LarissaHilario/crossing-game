package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	gameWidth  = 300
	gameHeight = 500
	playerSize = 50
	obstacleSize = 50
)

func main() {
	app := app.New()
	w := app.NewWindow("Crossy Road")
	w.Resize(fyne.NewSize(gameWidth, gameHeight))

	player := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 255})
	player.Resize(fyne.NewSize(playerSize, playerSize))
	player.Move(fyne.NewPos(float32(gameWidth/2-playerSize/2), float32(gameHeight-playerSize-10)))

	obstacle := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	obstacle.Resize(fyne.NewSize(obstacleSize, obstacleSize))
	obstacle.Move(fyne.NewPos(float32(rand.Intn(gameWidth-obstacleSize)), 0))

	score := 0
	scoreLabel := widget.NewLabel(fmt.Sprintf("Score: %d", score))

	gameObjects := container.NewWithoutLayout(player, obstacle, scoreLabel)
	w.SetContent(gameObjects)

	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * 100):
				player.Move(fyne.NewPos(player.Position().X, player.Position().Y-10))
				if player.Position().Y < 0 {
					player.Move(fyne.NewPos(player.Position().X, gameHeight-playerSize-10))
					score++
					scoreLabel.SetText(fmt.Sprintf("Score: %d", score))
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * 500):
				obstacle.Move(fyne.NewPos(float32(rand.Intn(gameWidth-obstacleSize)), 0))
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * 100):
				/*if player.Position().Intersects(obstacle.Position()) {
					fmt.Println("Game Over!")
					return
				}*/
			}
		}
	}()

	w.ShowAndRun()
}
