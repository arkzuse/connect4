package main

import (
	"testing"
	"github.com/google/uuid"
)

func TestGame(t *testing.T) {
	p1 := NewPlayer("player1")
	p2 := NewPlayer("player2")
	g := NewGame(p1.id, p2.id)

	// player order
	if g.players[0] != p1.id && g.players[1] != p2.id {
		t.Errorf("Expected players to be in order, got %v and %v", g.players[0], g.players[1])
	}

	// error first turn by second player
	if err := g.playTurn(p2.id, 0); err == nil {
		t.Errorf("Expected error for not player's turn, got nil")
	}

	// first turn by first player
	if err := g.playTurn(p1.id, 0); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// error second turn by first player
	if err := g.playTurn(p1.id, 0); err == nil {
		t.Errorf("Expected error for not player's turn, got nil")
	}

	// second turn by second player
	if err := g.playTurn(p2.id, 1); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	g.playTurn(p1.id, 0)
	g.playTurn(p2.id, 1)
	g.playTurn(p1.id, 0)
	g.playTurn(p2.id, 1)

	if g.checkWinner() != uuid.Nil {
		t.Errorf("Expected no winner, got %v", g.checkWinner())
	}

	g.playTurn(p1.id, 0)

	if g.checkWinner() != p1.id {
		t.Errorf("Expected winner to be player1, got %v", g.checkWinner())
	}
}