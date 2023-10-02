
package scenes

import (
	"fmt"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"juego-pollo/models"
	"juego-pollo/utils"
)
const (
    gameWidth     = 900
    gameHeight    = 500
    playerSize    = 50
    obstacleSize  = 50
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
        app:    app.New(),
        window: app.New().NewWindow("Crossy Road"),
        player: models.NewPlayer(),
        obstacle: models.NewObstacle(),
        scoreLabel: widget.NewLabel("Score: 0"),
		statusLabel: widget.NewLabel("vas bien"),
    }
	gameScene.statusLabel.Move(fyne.NewPos(50,460))
    gameScene.window.Resize(fyne.NewSize(gameWidth, gameHeight))
    gameScene.player.MoveTo(gameWidth/2-playerSize/2, gameHeight-playerSize-10)
    gameScene.gameObjects = container.NewWithoutLayout(gameScene.player.GetRectangle(), gameScene.obstacle.GetRectangle(), gameScene.scoreLabel, gameScene.statusLabel)
    gameScene.window.SetContent(gameScene.gameObjects)
    return gameScene
}




func (gs *GameScene) Start() {

    go func() {
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
    }()

    go func() {
        for {
            select {
            case <-time.After(time.Millisecond * 100):
		
                gs.obstacle.MoveTo(gs.obstacle.GetRectangle().Position().X+10, gs.obstacle.GetRectangle().Position().Y)
			  }
				
				if utils.ImageOverlaps(gs.player.GetRectangle(), gs.obstacle.GetRectangle()) {
			
				gs.statusLabel.SetText("Game Over!")
			
			return
            }
        }
    }()

	go func() {
		if utils.ImageOverlaps(gs.player.GetRectangle(), gs.obstacle.GetRectangle()) {
			gs.statusLabel.SetText("Game Over!")
			return
		}
			
		}()


    gs.window.ShowAndRun()
}