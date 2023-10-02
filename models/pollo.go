// models/player.go
package models

import (


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

const (

    gameHeight    = 500
    playerSize    = 50
  
)

type Player struct {
    rectangle *canvas.Image
    position  fyne.Position
}

func NewPlayer() *Player {
    player := &Player{
         rectangle:canvas.NewImageFromURI(storage.NewFileURI("./assets/pollo.png")),
    }
    player.rectangle.Resize(fyne.NewSize(playerSize, playerSize))
    return player
}

func (p *Player) MoveTo(x, y float32) {
    p.position = fyne.NewPos(x, y)
    p.rectangle.Move(p.position)
}

func (p *Player) GetRectangle() *canvas.Image {
    return p.rectangle
}
