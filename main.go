package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

func setupBaseGameInterface(gameManager *GameManager) *fyne.Container{
	gameBoard := gameManager.Board
	content := container.New(layout.NewGridLayout(gameManager.getBoardSize()))
	for i := range gameBoard.Board {
		for _, button := range gameBoard.Board[i]{
			content.Add(button)
		}
	}
	return content
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
