// scenes/game_scene.go
package scenes

import (
	"fmt"
	"math/rand"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"juego-pollo/models"
)
const (
    gameWidth     = 300
    gameHeight    = 500
    playerSize    = 50
    obstacleSize  = 70
)

type GameScene struct {
    app         fyne.App
    window      fyne.Window
    player      *models.Player
    obstacle    *models.Obstacle
    score       int
    scoreLabel  *widget.Label
    gameObjects fyne.CanvasObject
}

func NewGameScene() *GameScene {
    gameScene := &GameScene{
        app:    app.New(),
        window: app.New().NewWindow("Crossy Road"),
        player: models.NewPlayer(),
        obstacle: models.NewObstacle(),
        scoreLabel: widget.NewLabel("Score: 0"),
    }
    gameScene.window.Resize(fyne.NewSize(gameWidth, gameHeight))
    gameScene.player.MoveTo(gameWidth/2-playerSize/2, gameHeight-playerSize-10)
    gameScene.gameObjects = container.NewWithoutLayout(gameScene.player.GetRectangle(), gameScene.obstacle.GetRectangle(), gameScene.scoreLabel)
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
            case <-time.After(time.Millisecond * 250):
				
                gs.obstacle.MoveTo(float32(rand.Intn(gameWidth-obstacleSize)), 0)
				
            }
        }
    }()

   /* go func() {
        for {
            select {
            case <-time.After(time.Millisecond * 100):
                if gs.player.GetRectangle().Position().Intersects(gs.obstacle.GetRectangle().Position()) {
                    gs.GameOver()
                    return
                }
            }
        }
    }()*/

    gs.window.ShowAndRun()
}

func (gs *GameScene) GameOver() {
    // Implementar lógica para mostrar la pantalla de Game Over
}
