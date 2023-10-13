package scenes

import (
	"fmt"
	"juego-pollo/models"
	"juego-pollo/utils"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

const (
	gameWidth    = 300
	gameHeight   = 600
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
	moveUp      chan struct{}
	moveDown    chan struct{}
	moveLeft    chan struct{}
	moveRight   chan struct{}
}

func NewGameScene() *GameScene {
	gameScene := &GameScene{
		app:         app.New(),
		window:      app.New().NewWindow("Crossy Road"),
		player:      models.NewPlayer(),
		obstacle:    models.NewObstacle(),
		score:       3,
		scoreLabel:  widget.NewLabel("Vidas: 3 "),
		statusLabel: widget.NewLabel("vas bien"),
		moveUp:    make(chan struct{}),
		moveDown:  make(chan struct{}),
		moveLeft:  make(chan struct{}),
		moveRight: make(chan struct{}),
	}
	gameScene.statusLabel.Move(fyne.NewPos(160, 0))
	background := canvas.NewImageFromURI(storage.NewFileURI("./assets/paisaje.png"))
	background.Resize(fyne.NewSize(300, 500))
	background.Move(fyne.NewPos(0,0))
	pasto:= canvas.NewImageFromURI(storage.NewFileURI("./assets/pasto.png"))
	pasto.Resize(fyne.NewSize(300, 210))
	pasto.Move(fyne.NewPos(0,390))

	gameScene.window.Resize(fyne.NewSize(gameWidth, gameHeight))
	gameScene.player.MoveTo(gameWidth/2-playerSize/2, gameHeight-playerSize-10)
	gameScene.obstacle.MoveTo(gameWidth-obstacleSize+10, gameHeight/2-obstacleSize/2+130)
	gameScene.gameObjects = container.NewWithoutLayout(background,pasto,gameScene.player.GetRectangle(), gameScene.obstacle.GetRectangle(), gameScene.scoreLabel, gameScene.statusLabel)
	gameScene.window.SetContent(gameScene.gameObjects)

	
	
    
	return gameScene
}

	

	


func movePlayer(gs *GameScene) {
	for {

		select {
		case <-gs.moveUp:
			gs.player.MoveTo(gs.player.GetRectangle().Position().X, gs.player.GetRectangle().Position().Y-10)
			if gs.player.GetRectangle().Position().Y < 190 {
				gs.player.MoveTo(gs.player.GetRectangle().Position().X, 190)
			}
		case <-gs.moveDown:
			gs.player.MoveTo(gs.player.GetRectangle().Position().X, gs.player.GetRectangle().Position().Y+10)
			if gs.player.GetRectangle().Position().Y > gameHeight-playerSize {
				gs.player.MoveTo(gs.player.GetRectangle().Position().X, gameHeight-playerSize)
			}
		case <-gs.moveLeft:
			gs.player.MoveTo(gs.player.GetRectangle().Position().X-10, gs.player.GetRectangle().Position().Y)
			if gs.player.GetRectangle().Position().X < 0 {
				gs.player.MoveTo(0, gs.player.GetRectangle().Position().Y)
			}
		case <-gs.moveRight:
			gs.player.MoveTo(gs.player.GetRectangle().Position().X+10, gs.player.GetRectangle().Position().Y)
			if gs.player.GetRectangle().Position().X > gameWidth-playerSize {
				gs.player.MoveTo(gameWidth-playerSize, gs.player.GetRectangle().Position().Y)
			}
		}
		

	}
}

func moveObstacle(gs *GameScene) {
	for {
		select {
		case <-time.After(time.Millisecond * 20):
			gs.obstacle.MoveTo(gs.obstacle.GetRectangle().Position().X+5, gs.obstacle.GetRectangle().Position().Y)
			if gs.obstacle.GetRectangle().Position().X > gameWidth {
				gs.obstacle.MoveTo(-obstacleSize, gs.obstacle.GetRectangle().Position().Y)
			}
		}
	}
}

func checkCollision(gs *GameScene) {
	for {
		if utils.ImageOverlaps(gs.player.GetRectangle(), gs.obstacle.GetRectangle()) {
			gs.statusLabel.SetText("¡Ay te aplastaron!")
			gs.score--
			if gs.score == 0 {
				gs.statusLabel.SetText("Juego terminado")
				gs.window.Close()
				
			}
			gs.scoreLabel.SetText(fmt.Sprintf("Vidas: %d", gs.score))
			gs.player.MoveTo(gs.player.GetRectangle().Position().X, gameHeight-playerSize)
			
		}
		time.Sleep(time.Millisecond * 100)
		
	}
}



func (gs *GameScene) Start() {

	moveUpButton := widget.NewButton("Arriba", func() {
		gs.moveUp <- struct{}{}
	})
	moveDownButton := widget.NewButton("Abajo", func() {
		gs.moveDown <- struct{}{}
	})
	moveLeftButton := widget.NewButton("Izquierda", func() {
		gs.moveLeft <- struct{}{}
	})
	moveRightButton := widget.NewButton("Derecha", func() {
		gs.moveRight <- struct{}{}
	})
	
	// Mueve los botones a sus posiciones finales
	moveUpButton.Move(fyne.NewPos(10, 10))
	moveDownButton.Move(fyne.NewPos(10, 30))
	moveLeftButton.Move(fyne.NewPos(10, 50))
	moveRightButton.Move(fyne.NewPos(30, 50))

	// Cambia el tamaño de los botones
	moveUpButton.Resize(fyne.NewSize(50, 50))
	moveDownButton.Resize(fyne.NewSize(50, 50))
	moveLeftButton.Resize(fyne.NewSize(50, 50))
	moveRightButton.Resize(fyne.NewSize(50, 50))

	gs.window.SetContent(container.NewVBox(gs.gameObjects, moveUpButton, moveDownButton, moveLeftButton, moveRightButton, ))
	go movePlayer(gs)
	go moveObstacle(gs)
	go checkCollision(gs)
	// Muestra la ventana
	gs.window.ShowAndRun()

}
