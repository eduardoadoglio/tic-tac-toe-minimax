package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"math"
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
		GameWinner: "",
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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (gameManager *GameManager) minimax(gameBoard GameBoard, depth int, isMaximizing bool) int {
	scores := map[string]int {
		"X": -1 - (1/1+depth),
		"O": 1 - (1/1+depth),
		"TIE": 0,
	}
	if gameManager.isGameOver() {
		return scores[gameManager.getWinner()]
	}
	bestScore := math.MaxInt64
	if isMaximizing {
		bestScore = math.MinInt64
	}
	for i := range gameBoard.Board {
		for j := range gameBoard.Board[i] {
			if gameBoard.Board[i][j].Text == "" {
				if isMaximizing {
					gameBoard.Board[i][j].Text = "O"
				}else {
					gameBoard.Board[i][j].Text = "X"
				}
				score := gameManager.minimax(gameBoard, depth + 1, !isMaximizing)
				gameBoard.Board[i][j].Text = ""
				if isMaximizing {
					bestScore = max(score, bestScore)
				} else {
					bestScore = min(score, bestScore)
				}
			}
		}
	}
	return bestScore
}

func (gameManager *GameManager) handleAiMove() {
	fmt.Println("-- Handling AI turn")
	gameBoard := gameManager.Board
	bestScore := math.MinInt64
	bestMove := [2]int {}
	for i := range gameBoard.Board {
		for j, button := range gameBoard.Board[i] {
			if button.Text == "" {
				button.Text = "O"
				score := gameManager.minimax(*gameBoard, 0, false)
				button.Text = ""
				if score > bestScore {
					bestScore = score
					bestMove[0] = i
					bestMove[1] = j
				}
			}
		}
	}
	fmt.Printf("-- Setting button [%d][%d]\n", bestMove[0], bestMove[1])
	gameBoard.SetText("O", bestMove[0], bestMove[1])
	if gameManager.isGameOver() {
		gameManager.handleGameOver()
		fmt.Printf("%s won.\n", gameManager.GameWinner)
	}
	gameManager.setCurrentPlayer("X")
}

func (gameManager *GameManager) checkHorizontalWinner() (bool, string) {
	gameBoard := gameManager.Board
	for i := range gameBoard.Board {
		xWonHorizontal := true
		oWonHorizontal := true
		for _, button := range gameBoard.Board[i] {
			if button.Text != "X" {
				xWonHorizontal = false
			}
			if button.Text != "O" {
				oWonHorizontal = false
			}
		}
		if xWonHorizontal {
			return true, "X"
		} else if oWonHorizontal {
			return true, "O"
		}
	}
	return false, ""
}

func (gameManager *GameManager) checkVerticalWinner() (bool, string) {
	gameBoard := gameManager.Board
	for i := range gameBoard.Board {
		xWonVertical := true
		oWonVertical := true
		for j := range gameBoard.Board[i] {
			if gameBoard.Board[j][i].Text != "X" {
				xWonVertical = false
			}
			if gameBoard.Board[j][i].Text != "O" {
				oWonVertical = false
			}
		}
		if xWonVertical {
			return true, "X"
		}else if oWonVertical {
			return true, "O"
		}
	}
	return false, ""
}

func (gameManager *GameManager) checkDiagonalWinner() (bool, string) {
	gameBoard := gameManager.Board
	diagonals := [][]*widget.Button {
		{gameBoard.Board[0][0], gameBoard.Board[1][1], gameBoard.Board[2][2]},
		{gameBoard.Board[0][2], gameBoard.Board[1][1], gameBoard.Board[2][0]},
	}

	for i := range diagonals {
		xWonDiagonal := true
		oWonDiagonal := true
		for _, button := range diagonals[i] {
			if button.Text != "X" {
				xWonDiagonal = false
			}
			if button.Text != "O" {
				oWonDiagonal = false
			}
		}
		if xWonDiagonal {
			return true, "X"
		} else if oWonDiagonal {
			return true, "O"
		}
	}
	return false, ""
}

func (gameManager *GameManager) checkForTies() (bool, string) {
	gameBoard := gameManager.Board
	for i := range gameBoard.Board {
		for j := range gameBoard.Board[i] {
			if gameBoard.Board[j][i].Text == "" {
				return false, ""
			}
		}
	}
	gameManager.CurrentPlayer = "TIE"
	return true, "TIE"
}

func (gameManager *GameManager) handleGameOver() {
	gameManager.GameWinner = gameManager.CurrentPlayer
	gameManager.GameState = "OVER"
}

func (gameManager *GameManager) isGameOver() bool {
	horizontalWinner, _ := gameManager.checkHorizontalWinner()
	verticalWinner, _ := gameManager.checkVerticalWinner()
	diagonalWinner, _ := gameManager.checkDiagonalWinner()
	isTied, _ := gameManager.checkForTies()
	return horizontalWinner || verticalWinner || diagonalWinner || isTied
}

func (gameManager *GameManager) getWinner() string {
	horizontalWinner, winner := gameManager.checkHorizontalWinner()
	if horizontalWinner {
		return winner
	}
	verticalWinner, winner := gameManager.checkVerticalWinner()
	if verticalWinner {
		return winner
	}
	diagonalWinner, winner := gameManager.checkDiagonalWinner()
	if diagonalWinner {
		return winner
	}
	isTied, winner := gameManager.checkForTies()
	if isTied {
		return winner
	}
	return ""
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
				if gameManager.isGameOver() {
					gameManager.handleGameOver()
					fmt.Printf("%s won.\n", gameManager.getWinner())
					return
				}
				gameManager.setCurrentPlayer("O")
				gameManager.handleAiMove()
			}
		}
	}
}