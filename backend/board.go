package main

import "fmt"

const (
	rows = 6
	cols = 7
)

type Board struct {
	Board [rows][cols]int
}

func NewBoard() *Board {
	var b Board

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			b.Board[i][j] = -1
		}
	}

	return &b
}

func (b *Board) placePiece(col int, piece int) error {
	if col < 0 || col >= cols {
		return fmt.Errorf("invalid column")
	}

	if (piece < 0 || piece >= 2) {
		return fmt.Errorf("invalid piece")
	}

	for i := rows - 1; i >= 0; i-- {
		if b.Board[i][col] == -1 {
			b.Board[i][col] = piece
			return nil
		}
	}

	return fmt.Errorf("column %d is full", col)
}

func (b *Board) checkBoard() int {
	dirs := [3][2]int{{-1, 0}, {0, -1}, {-1, -1}}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if b.Board[i][j] == -1 {
				continue
			}

			for _, dir := range dirs {
				if i+3*dir[0] < 0 || j+3*dir[1] < 0 {
					continue
				}

				cur := b.Board[i][j]

				for k := 1; k < 4; k++ {
					if b.Board[i+k*dir[0]][j+k*dir[1]] != cur {
						break
					}

					if k == 3 {
						return cur
					}
				}
			}
		}
	}

	return -1
}
