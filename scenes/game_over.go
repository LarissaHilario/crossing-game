package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
    window fyne.Window
}

func NewGameOverScene() *GameOverScene {
    return &GameOverScene{
     
        window: app.New().NewWindow("Juego terminado"),
    }
}

func (gos *GameOverScene) Show() {
    gameOverLabel := widget.NewLabel("Juego terminado")
    gos.window.SetContent(container.NewVBox(gameOverLabel))
    gos.window.Show()
}
