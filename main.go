package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var gameBoard = [][] string {
	{"", "", ""},
	{"", "", ""},
	{"", "", ""},
}

var buttonBoard = make([][]*widget.Button, len(gameBoard))

func initButtonBoard(){
	for i := range buttonBoard {
		buttonBoard[i] = make([]*widget.Button, len(gameBoard))
	}
}

func handleCurrentTurn(row, col int) func() {
	return func(){
		buttonBoard[row][col].SetText("X")
		gameBoard[row][col] = "X"
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tic-tac-toe")
	myApp.Settings().SetTheme(theme.LightTheme())

	myWindow.Resize(fyne.NewSize(400, 400))

	content := container.New(layout.NewGridLayout(len(gameBoard)))
	initButtonBoard()
	for i := 0; i < len(gameBoard); i++ {
		for j := 0; j < len(gameBoard[i]); j++ {
			button := widget.NewButton(gameBoard[i][j], handleCurrentTurn(i, j))
			content.Add(button)
			buttonBoard[i][j] = button
		}
	}
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
