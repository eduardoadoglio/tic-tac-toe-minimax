package main

import "fyne.io/fyne/v2/widget"

type Board [][]*widget.Button

func (board *Board) NewBoard(size int) *Board {
	newBoard := make(Board, size)
	for i := range newBoard {
		newBoard[i] = make([]*widget.Button, size)
	}
	newBoard.initBoard()
	return &newBoard
}

func (board *Board) initBoard() *Board{
	for i := range *board {
		for j := range (*board)[i] {
			button := widget.NewButton("", HandleCurrentTurn(i, j))
			board.SetButtonAtIndex(button, i, j)
		}
	}
	return board
}

func (board *Board) SetButtonAtIndex(button *widget.Button, row, col int) *Board{
	(*board)[row][col] = button
	return board
}

func (board *Board) SetText(value string, row, col int) *Board {
	(*board)[row][col].SetText(value)
	return board
}

func (board *Board) GetText(row, col int) string {
	return (*board)[row][col].Text
}

func (board *Board) ResetBoard() *Board {
	return board.NewBoard(len(*(board)))
}