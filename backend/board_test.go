package main

import (
	"testing"
)

func TestPlacePiece(t *testing.T) {
	b := newBoard()

	err := b.placePiece(0, "X")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = b.placePiece(0, "O")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	for i := 0; i < 4; i++ {
		err = b.placePiece(1, "X")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}

	w := b.checkBoard()
	if w != "X" {
		t.Errorf("Expected X to win, got %s", w)
	}

	/// test diagonal win
	b = newBoard()

	for i := 0; i < 4; i++ {
		b.Board[i][i] = "X"
	}

	w = b.checkBoard()
	if w != "X" {
		t.Errorf("Expected X to win, got %s", w)
	}
}