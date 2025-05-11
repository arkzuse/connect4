package main

import (
	"testing"
)

func TestPlacePiece(t *testing.T) {
	b := NewBoard()

	err := b.placePiece(0, 0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = b.placePiece(0, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	for i := 0; i < 4; i++ {
		err = b.placePiece(1, 0)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}

	w := b.checkBoard()
	if w != 0 {
		t.Errorf("Expected X to win, got %d", w)
	}

	/// test diagonal win
	b = NewBoard()

	for i := 0; i < 4; i++ {
		b.Board[i][i] = 0
	}

	w = b.checkBoard()
	if w != 0 {
		t.Errorf("Expected X to win, got %d", w)
	}
}
