package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)
const (
    obstacleSize = 70
    gameWidth    = 300
)

type Obstacle struct {
    rectangle *canvas.Image
    position  fyne.Position
}

func NewObstacle() *Obstacle {
    obstacle := &Obstacle{
        rectangle: canvas.NewImageFromURI(storage.NewFileURI("./assets/tracktor.png")),
    }
    obstacle.rectangle.Resize(fyne.NewSize(obstacleSize, obstacleSize))
    
    return obstacle
}

func (o *Obstacle) MoveTo(x, y float32) {
    o.position = fyne.NewPos(x, y)
    o.rectangle.Move(o.position)
}

func (o *Obstacle) GetRectangle() *canvas.Image {
    return o.rectangle
}
