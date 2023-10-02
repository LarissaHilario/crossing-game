package models

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Pollo struct {
    posX, posY  float32
    direction   int // 0 para parado, -1 para izquierda, 1 para derecha
    status      bool
    pel         *canvas.Image
}


func NewPollo(posx float32, posy float32, img *canvas.Image) *Pollo {
	return &Pollo{
		posX: posx,
		posY: posy,
		status: true,
		pel: img,
	}
}

func (p *Pollo) Run() {
    var stepSize float32 = 5.0 // Tamaño de paso para el movimiento

    p.status = true
    for p.status {
        if p.direction == -1 { // Mover a la izquierda
            if p.posX-stepSize >= 0 {
                p.posX -= stepSize
            }
        } else if p.direction == 1 { // Mover a la derecha
            if p.posX+stepSize <= 740 { // 740 es el límite derecho
                p.posX += stepSize
            }
        }
        p.pel.Move(fyne.NewPos(p.posX, p.posY))
        time.Sleep(100 * time.Millisecond)
    }
}
