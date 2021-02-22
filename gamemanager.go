package main

import "fmt"

type GameManager struct {
	Board *GameBoard
	CurrentPlayer string
}

func (gameManager *GameManager) getBoardSize() int {
	return gameManager.Board.Size
}

func NewGameManager(boardSize int) *GameManager{
	fmt.Println("-- Creating new GameManager object")
	gameManager := &GameManager{
		CurrentPlayer: "X",
	}
	gameManager.Board = NewBoard(boardSize, gameManager.HandleCurrentTurn)
	return gameManager
}

func (gameManager *GameManager) setCurrentPlayer(currentPlayer string) {
	gameManager.CurrentPlayer = currentPlayer
}

func (gameManager *GameManager) HandleCurrentTurn(row, col int) func() {
	return func(){
		gameManager.Board.SetText("X", row, col)
	}
}