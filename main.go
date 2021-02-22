package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func setupBaseGameInterface(gameManager *GameManager) *fyne.Container {
	gameBoard := gameManager.Board
	gameGrid := container.New(layout.NewGridLayout(gameManager.getBoardSize()))
	for i := range gameBoard.Board {
		for _, button := range gameBoard.Board[i]{
			gameGrid.Add(button)
		}
	}
	resetButton := container.NewCenter(
		widget.NewButtonWithIcon("", theme.MediaReplayIcon(), gameManager.ResetGame),
	)
	gameGrid.Add(layout.NewSpacer())
	gameGrid.Add(resetButton)
	gameGrid.Add(layout.NewSpacer())
	return gameGrid
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tic-tac-toe")
	myApp.Settings().SetTheme(theme.LightTheme())
	gameManager := NewGameManager(3)
	myWindow.Resize(fyne.NewSize(400, 400))
	baseInterface := setupBaseGameInterface(gameManager)
	myWindow.SetContent(baseInterface)
	myWindow.ShowAndRun()
}
