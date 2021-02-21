package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

var gameBoard *Board

func setupBaseGameInterface() *fyne.Container{
	content := container.New(layout.NewGridLayout(len(*gameBoard)))
	for i := range *gameBoard {
		for _, button := range (*gameBoard)[i]{
			content.Add(button)
		}
	}
	return content
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tic-tac-toe")
	myApp.Settings().SetTheme(theme.LightTheme())
	gameBoard = gameBoard.NewBoard(3)
	myWindow.Resize(fyne.NewSize(400, 400))
	baseInterface := setupBaseGameInterface()
	myWindow.SetContent(baseInterface)
	myWindow.ShowAndRun()
}
