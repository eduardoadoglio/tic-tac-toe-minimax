package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func appendActionButtons(gameGrid *fyne.Container, gameManager *GameManager) {
	resetButton := container.NewCenter(
		widget.NewButtonWithIcon("", theme.MediaReplayIcon(), gameManager.ResetGame),
	)
	gameGrid.Add(layout.NewSpacer())
	gameGrid.Add(resetButton)
	gameGrid.Add(layout.NewSpacer())
}

func createGameGrid(gameManager *GameManager) *fyne.Container {
	gameBoard := gameManager.Board
	gameGrid := container.New(layout.NewGridLayout(gameManager.getBoardSize()))
	for i := range gameBoard.Board {
		for _, button := range gameBoard.Board[i]{
			gameGrid.Add(button)
		}
	}
	return gameGrid
}

func setupBaseGameInterface(gameManager *GameManager) *fyne.Container {
	gameGrid := createGameGrid(gameManager)
	appendActionButtons(gameGrid, gameManager)
	return gameGrid
}

func setupMenuInterface() *fyne.Container {
	menu := container.New(layout.NewGridLayoutWithRows(3))
	title := container.NewCenter(widget.NewLabel("Selecione o jogador"))
	xButton := widget.NewButton("X", nil)
	oButton := widget.NewButton("O", nil)
	menu.Add(title)
	menu.Add(xButton)
	menu.Add(oButton)
	return menu
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Tic-tac-toe")
	myApp.Settings().SetTheme(theme.LightTheme())
	// players := NewPlayers("X", "O")
	// gameManager := NewGameManager(3, players)
	myWindow.Resize(fyne.NewSize(400, 400))
	// baseInterface := setupBaseGameInterface(gameManager)
	baseInterface := setupMenuInterface()
	myWindow.SetContent(baseInterface)
	myWindow.ShowAndRun()
}
