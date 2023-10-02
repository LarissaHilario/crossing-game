package scenes

import (
	"fmt"
	"juego-pollo/models"
	"juego-pollo/utils"
	
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	gameWidth    = 300
	gameHeight   = 500
	playerSize   = 50
	obstacleSize = 50
	
)

type GameScene struct {
	app         fyne.App
	window      fyne.Window
	player      *models.Player
	obstacle    *models.Obstacle
	score       int
	scoreLabel  *widget.Label
	statusLabel *widget.Label
	gameObjects fyne.CanvasObject
	
}

func NewGameScene() *GameScene {
	gameScene := &GameScene{
		app:         app.New(),
		window:      app.New().NewWindow("Crossy Road"),
		player:      models.NewPlayer(),
		obstacle:    models.NewObstacle(),
		scoreLabel:  widget.NewLabel("Score: 0"),
		statusLabel: widget.NewLabel("vas bien"),
	}
	gameScene.statusLabel.Move(fyne.NewPos(50, 460))
	gameScene.window.Resize(fyne.NewSize(gameWidth, gameHeight))
	gameScene.player.MoveTo(gameWidth/2-playerSize/2, gameHeight-playerSize-10)

	gameScene.gameObjects = container.NewWithoutLayout(gameScene.player.GetRectangle(), gameScene.obstacle.GetRectangle(), gameScene.scoreLabel, gameScene.statusLabel)
	gameScene.window.SetContent(gameScene.gameObjects)
	return gameScene
}

func movePlayer(gs *GameScene) {
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			gs.player.MoveTo(gs.player.GetRectangle().Position().X, gs.player.GetRectangle().Position().Y-10)
			if gs.player.GetRectangle().Position().Y < 0 {
				gs.player.MoveTo(gs.player.GetRectangle().Position().X, gameHeight-playerSize-10)
				gs.score++
				gs.scoreLabel.SetText(fmt.Sprintf("Score: %d", gs.score))
			}
		}
	}
}

func moveObstacle(gs *GameScene) {
    for {
        select {
        case <-time.After(time.Millisecond * 100):
            gs.obstacle.MoveTo(gs.obstacle.GetRectangle().Position().X+10, gs.obstacle.GetRectangle().Position().Y)
            if gs.obstacle.GetRectangle().Position().X > gameWidth {
                gs.obstacle.MoveTo(-obstacleSize, gs.obstacle.GetRectangle().Position().Y)
            }
        }
    }
}


func checkCollision(gs *GameScene) {
	for {
		select {
		case <-time.After(time.Millisecond * 100):
			if utils.ImageOverlaps(gs.player.GetRectangle(), gs.obstacle.GetRectangle()) {
				gs.statusLabel.SetText("¡Ay te aplastaron!")
				return
			}
		}
	}
}

func (gs *GameScene) Start() {
	
	go movePlayer(gs)
	go moveObstacle(gs)
	go checkCollision(gs)
	gs.window.ShowAndRun()
}
