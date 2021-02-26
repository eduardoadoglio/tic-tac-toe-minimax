package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
)

type GameBoard struct {
	Board [][]*widget.Button
	Size int
	OnPressed func (int, int) func()
}

func NewBoard(size int, onPressed func(int, int) func()) *GameBoard {
	fmt.Println("-- Creating new board object")
	gameBoard := &GameBoard{
		Size: size,
		OnPressed: onPressed,
	}
	gameBoard.Board = gameBoard.createEmptyBoard()
	return gameBoard
}

func (gameBoard *GameBoard) createEmptyBoard() [][]*widget.Button{
	fmt.Println("-- Creating empty board")
	newBoard := make([][]*widget.Button, gameBoard.Size)
	for i := range newBoard {
		newBoard[i] = make([]*widget.Button, gameBoard.Size)
	}
	gameBoard.Board = newBoard
	return gameBoard.initBoard()
}

func (gameBoard *GameBoard) initBoard() [][]*widget.Button{
	fmt.Println("-- Initializing empty board")
	for i := range (*gameBoard).Board {
		for j := range (*gameBoard).Board[i] {
			button := widget.NewButton("", (*gameBoard).OnPressed(i, j))
			gameBoard.SetButtonAtIndex(button, i, j)
		}
	}
	return gameBoard.Board
}

func (gameBoard *GameBoard) SetButtonAtIndex(button *widget.Button, row, col int) *GameBoard{
	(*gameBoard).Board[row][col] = button
	return gameBoard
}

func (gameBoard *GameBoard) SetText(value string, row, col int) {
	(*gameBoard).Board[row][col].SetText(value)
}

func (gameBoard *GameBoard) GetText(row, col int) string {
	return (*gameBoard).Board[row][col].Text
}

func (gameBoard *GameBoard) ResetBoard() {
	for i := range gameBoard.Board {
		for _, button := range gameBoard.Board[i] {
			button.SetText("")
		}
	}
}

func (gameBoard *GameBoard) printGameBoard() {
	for i := range gameBoard.Board {
		for _, button := range gameBoard.Board[i] {
			fmt.Printf("\"%s\" ", button.Text)
		}
		fmt.Println()
	}
}