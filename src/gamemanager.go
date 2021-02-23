package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
)

type GameManager struct {
	Board *GameBoard
	CurrentPlayer string
	GameState string
	GameWinner string
}

func NewGameManager(boardSize int) *GameManager{
	fmt.Println("-- Creating new GameManager object")
	gameManager := &GameManager{
		CurrentPlayer: "X",
		GameState: "IN_PROGRESS",
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
				if gameManager.checkWinner() {
					fmt.Printf("%s won.\n", gameManager.GameWinner)
				}
				gameManager.setCurrentPlayer("X")
				return
			}
		}
	}
}

func (gameManager *GameManager) checkHorizontalWinner() bool {
	gameBoard := gameManager.Board
	for i := range gameBoard.Board {
		wonHorizontal := true
		for _, button := range gameBoard.Board[i] {
			if button.Text != gameManager.CurrentPlayer {
				wonHorizontal = false
			}
		}
		if wonHorizontal {
			return wonHorizontal
		}
	}
	return false
}

func (gameManager *GameManager) checkVerticalWinner() bool {
	gameBoard := gameManager.Board
	for i := range gameBoard.Board {
		wonVertical := true
		for j := range gameBoard.Board[i] {
			if gameBoard.Board[j][i].Text != gameManager.CurrentPlayer {
				wonVertical = false
			}
		}
		if wonVertical {
			return wonVertical
		}
	}
	return false
}

func (gameManager *GameManager) checkDiagonalWinner() bool {
	gameBoard := gameManager.Board
	diagonals := [][]*widget.Button {
		{gameBoard.Board[0][0], gameBoard.Board[1][1], gameBoard.Board[2][2]},
		{gameBoard.Board[0][2], gameBoard.Board[1][1], gameBoard.Board[2][0]},
	}

	for i := range diagonals {
		wonDiagonal := true
		for _, button := range diagonals[i] {
			if button.Text != gameManager.CurrentPlayer {
				wonDiagonal = false
			}
		}
		if wonDiagonal {
			return wonDiagonal
		}
	}
	return false
}


func (gameManager *GameManager) checkWinner() bool {
	horizontalWinner := gameManager.checkHorizontalWinner()
	verticalWinner := gameManager.checkVerticalWinner()
	diagonalWinner := gameManager.checkDiagonalWinner()

	if horizontalWinner || verticalWinner || diagonalWinner {
		gameManager.GameWinner = gameManager.CurrentPlayer
		gameManager.GameState = "OVER"
		return true
	}

	return false
}

func (gameManager *GameManager) ResetGame() {
	gameBoard := gameManager.Board
	gameBoard.ResetBoard()
	gameManager.setCurrentPlayer("X")
	gameManager.GameWinner = ""
	gameManager.GameState = "IN_PROGRESS"
}

func (gameManager *GameManager) HandleCurrentTurn(row, col int) func() {
	return func(){
		if gameManager.GameState == "IN_PROGRESS" {
			gameBoard := gameManager.Board
			if gameManager.CurrentPlayer == "X" && gameBoard.GetText(row, col) == "" {
				gameBoard.SetText("X", row, col)
				if gameManager.checkWinner() {
					fmt.Printf("%s won.\n", gameManager.GameWinner)
					return
				}
				gameManager.setCurrentPlayer("O")
				gameManager.handleAiMove()
			}
		}
	}
}