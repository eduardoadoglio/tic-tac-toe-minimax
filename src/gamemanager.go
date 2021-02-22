package main

import "fmt"

type GameManager struct {
	Board *GameBoard
	CurrentPlayer string
}

func NewGameManager(boardSize int) *GameManager{
	fmt.Println("-- Creating new GameManager object")
	gameManager := &GameManager{
		CurrentPlayer: "X",
	}
	gameManager.Board = NewBoard(boardSize, gameManager.HandleCurrentTurn)
	return gameManager
}

func (gameManager *GameManager) getBoardSize() int {
	return gameManager.Board.Size
}

func (gameManager *GameManager) setCurrentPlayer(currentPlayer string) {
	gameManager.CurrentPlayer = currentPlayer
}

func (gameManager *GameManager) handleAiMove() {
	gameBoard := gameManager.Board
	for i := range gameBoard.Board {
		for j, button := range gameBoard.Board[i] {
			if button.Text == "" {
				gameBoard.SetText("O", i, j)
				gameManager.setCurrentPlayer("X")
				return
			}
		}
	}
}

func (gameManager *GameManager) ResetGame() {
	gameBoard := gameManager.Board
	gameBoard.ResetBoard()
	gameManager.setCurrentPlayer("X")
}

func (gameManager *GameManager) HandleCurrentTurn(row, col int) func() {
	return func(){
		fmt.Printf("-- Player clicked on [%d][%d]\n", row, col)
		gameBoard := gameManager.Board
		if gameManager.CurrentPlayer == "X" && gameBoard.GetText(row, col) == "" {
			fmt.Printf("-- Player valid move on [%d][%d]\n", row, col)
			gameBoard.SetText("X", row, col)
			gameManager.setCurrentPlayer("O")
			gameManager.handleAiMove()
		}
	}
}