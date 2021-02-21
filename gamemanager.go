package main

import "fyne.io/fyne/v2/widget"

type GameManager struct {
	Board Board
	ButtonBoard [][]*widget.Button
	CurrentPlayer string
}

func (gameManager *GameManager) NewGameManager() *GameManager{
	// placeholder
	return gameManager
}

func HandleCurrentTurn(row, col int) func() {
	return func(){
		gameBoard.SetText("X", row, col)
	}
}