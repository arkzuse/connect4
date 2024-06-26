package main

import "fmt"

const (
	rows = 6
	cols = 7
)

type Board struct {
	Board [rows][cols]string
}

func newBoard() *Board {
	var b Board

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			b.Board[i][j] = ""
		}
	}

	return &b
}

func (b *Board) placePiece(col int, player string) error {
	for i := rows - 1; i >= 0; i-- {
		if b.Board[i][col] == "" {
			b.Board[i][col] = player
			return nil
		}
	}

	return fmt.Errorf("column %d is full", col)
}

func (b *Board) checkBoard() string {
	dirs := [3][2]int{{-1, 0}, {0, -1}, {-1, -1}}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if b.Board[i][j] == "" {
				continue
			}

			for _, dir := range dirs {
				if i + 3 * dir[0] < 0 || j + 3 * dir[1] < 0 {
					continue
				}

				cur := b.Board[i][j]

				for k := 1; k < 4; k++ {
					if b.Board[i + k * dir[0]][j + k * dir[1]] != cur {
						break
					}

					if k == 3 {
						return cur
					}
				}
			}
		}
	} 

	return ""
}