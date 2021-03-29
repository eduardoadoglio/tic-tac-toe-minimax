package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	app fyne.App
	window fyne.Window
	gameManager *GameManager
}

func NewGUI() *GUI {
	gui := &GUI{
		app: app.New(),
	}
	gui.window = gui.app.NewWindow("Tic-tac-toe")
	gui.app.Settings().SetTheme(theme.LightTheme())
	gui.window.Resize(fyne.NewSize(400, 400))
	return gui
}

func (gui *GUI) setWindowContent(content fyne.CanvasObject) {
	gui.window.SetContent(content)
}

func (gui *GUI) showContentAndRun() {
	gui.window.ShowAndRun()
}

func (gui *GUI) appendActionButtons(gameGrid *fyne.Container) {
	quitButton := container.NewCenter(
		widget.NewButtonWithIcon("", theme.CancelIcon(), gui.quitGame),
	)
	resetButton := container.NewCenter(
		widget.NewButtonWithIcon("", theme.MediaReplayIcon(), gui.gameManager.ResetGame),
	)
	returnButton := container.NewCenter(
		widget.NewButtonWithIcon("", theme.ContentUndoIcon(), gui.returnToMenu),
	)
	gameGrid.Add(quitButton)
	gameGrid.Add(resetButton)
	gameGrid.Add(returnButton)
}

func (gui *GUI) prependWinIndicator(gameGrid *fyne.Container) {
	gameGrid.Add(layout.NewSpacer())
	winIndicator := widget.NewLabel("")
	gui.gameManager.WinIndicator = winIndicator
	gameGrid.Add(container.NewCenter(winIndicator))
	gameGrid.Add(layout.NewSpacer())
}

func (gui *GUI) createGameGrid() *fyne.Container {
	gameManager := gui.gameManager
	gameBoard := gameManager.Board
	gameGrid := container.New(layout.NewGridLayout(gameManager.getBoardSize()))
	gui.prependWinIndicator(gameGrid)
	for i := range gameBoard.Board {
		for _, button := range gameBoard.Board[i]{
			gameGrid.Add(button)
		}
	}
	return gameGrid
}

func (gui *GUI) setupBaseGameInterface() *fyne.Container {
	gameGrid := gui.createGameGrid()
	gui.appendActionButtons(gameGrid)
	return gameGrid
}

func (gui *GUI) setupMenuInterface() *fyne.Container {
	menu := container.New(layout.NewGridLayoutWithRows(3))
	title := container.NewCenter(widget.NewLabel("Selecione o jogador"))
	xButton := widget.NewButton("X", gui.initGameInterface("X"))
	oButton := widget.NewButton("O", gui.initGameInterface("O"))
	menu.Add(title)
	menu.Add(xButton)
	menu.Add(oButton)
	return menu
}

func (gui *GUI) initGameInterface(humanPlayer string) func() {
	return func() {
		gui.gameManager = NewGameManager(3, humanPlayer)
		gameInterface := gui.setupBaseGameInterface()
		gui.setWindowContent(gameInterface)
	}
}

func (gui *GUI) quitGame() {
	gui.window.Close()
}

func (gui *GUI) returnToMenu() {
	baseInterface := gui.setupMenuInterface()
	gui.setWindowContent(baseInterface)
}