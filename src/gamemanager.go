package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"math"
)

type GameManager struct {
	Board *GameBoard
	Players *Players
	CurrentPlayer string
	GameState string
	GameWinner string
}

func NewGameManager(boardSize int, humanPlayer string) *GameManager{
	fmt.Println("-- Creating new GameManager object")
	gameManager := &GameManager{
		CurrentPlayer: "X",
		GameWinner: "",
		GameState: "IN_PROGRESS",
	}
	gameManager.setHumanPlayer(humanPlayer)
	gameManager.Board = NewBoard(boardSize, gameManager.handleHumanTurn)
	if gameManager.CurrentPlayer == gameManager.Players.AI {
		gameManager.handleAiTurn()
	}
	return gameManager
}

func (gameManager *GameManager) getBoardSize() int {
	return gameManager.Board.Size
}

func (gameManager *GameManager) setCurrentPlayer(currentPlayer string) {
	gameManager.CurrentPlayer = currentPlayer
}

func (gameManager *GameManager) setHumanPlayer(humanPlayer string) {
	players := NewPlayersWithHuman(humanPlayer)
	gameManager.Players = players
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
		gameManager.Players.Human: -1,
		gameManager.Players.AI: 1,
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
					gameBoard.Board[i][j].Text = gameManager.Players.AI
				}else {
					gameBoard.Board[i][j].Text = gameManager.Players.Human
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

func (gameManager *GameManager) handleAiTurn() {
	fmt.Println("-- Handling AI turn")
	gameBoard := gameManager.Board
	bestScore := math.MinInt64
	bestMove := [2]int {}
	for i := range gameBoard.Board {
		for j, button := range gameBoard.Board[i] {
			if button.Text == "" {
				button.Text = gameManager.Players.AI
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
	gameBoard.SetText(gameManager.Players.AI, bestMove[0], bestMove[1])
	if gameManager.isGameOver() {
		gameManager.handleGameOver()
		fmt.Printf("%s won.\n", gameManager.GameWinner)
	}
	gameManager.setCurrentPlayer(gameManager.Players.Human)
}

func (gameManager *GameManager) setGameWinner(gameWinner string) {
	gameManager.GameWinner = gameWinner
}

func (gameManager *GameManager) checkHorizontalWinner() bool {
	gameBoard := gameManager.Board
	currentPlayer := gameManager.CurrentPlayer
	for i := range gameBoard.Board {
		currentPlayerWonHorizontal := true
		for _, button := range gameBoard.Board[i] {
			if button.Text != currentPlayer {
				currentPlayerWonHorizontal = false
			}
		}
		if currentPlayerWonHorizontal {
			gameManager.setGameWinner(currentPlayer)
			return true
		}
	}
	return false
}

func (gameManager *GameManager) checkVerticalWinner() bool {
	gameBoard := gameManager.Board
	currentPlayer := gameManager.CurrentPlayer
	for i := range gameBoard.Board {
		currentPlayerWonVertical := true
		for j := range gameBoard.Board[i] {
			if gameBoard.Board[j][i].Text != currentPlayer {
				currentPlayerWonVertical = false
			}
		}
		if currentPlayerWonVertical {
			gameManager.setGameWinner(currentPlayer)
			return true
		}
	}
	return false
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
	return true, "TIE"
}

func (gameManager *GameManager) paintWinner() {
	return
}

func (gameManager *GameManager) handleGameOver() {
	gameManager.GameWinner = gameManager.CurrentPlayer
	gameManager.GameState = "OVER"
	gameManager.paintWinner()
}

func (gameManager *GameManager) isGameOver() bool {
	horizontalWinner := gameManager.checkHorizontalWinner()
	verticalWinner := gameManager.checkVerticalWinner()
	diagonalWinner, _ := gameManager.checkDiagonalWinner()
	isTied, _ := gameManager.checkForTies()
	return horizontalWinner || verticalWinner || diagonalWinner || isTied
}

func (gameManager *GameManager) getWinner() string {
	horizontalWinner := gameManager.checkHorizontalWinner()
	verticalWinner := gameManager.checkVerticalWinner()
	if horizontalWinner || verticalWinner {
		return gameManager.GameWinner
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
	if gameManager.CurrentPlayer == gameManager.Players.AI {
		gameManager.handleAiTurn()
	}
}

func (gameManager *GameManager) handleHumanTurn(row, col int) func() {
	return func(){
		if gameManager.GameState == "IN_PROGRESS" {
			gameBoard := gameManager.Board
			if gameManager.CurrentPlayer == gameManager.Players.Human && gameBoard.GetText(row, col) == "" {
				gameBoard.SetText(gameManager.Players.Human, row, col)
				if gameManager.isGameOver() {
					gameManager.handleGameOver()
					fmt.Printf("%s won.\n", gameManager.getWinner())
					return
				}
				gameManager.setCurrentPlayer(gameManager.Players.AI)
				gameManager.handleAiTurn()
			}
		}
	}
}