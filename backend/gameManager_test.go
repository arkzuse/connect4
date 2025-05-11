package main

import (
	"testing"
)

func TestGameManager(t * testing.T) {
	gm := NewGameManager()

	p1 := NewPlayer("player1")
	p2 := NewPlayer("player2")
	g1 := gm.newGame(p1.id, p2.id)

	if gm.getGameById(g1.id) != g1 {
		t.Errorf("Expected game %v, got %v", g1, gm.games[g1.id])
	}

	gm.removeGame(g1.id)
	if gm.getGameById(g1.id) != nil {
		t.Errorf("Expecter player to be nil, got %v", gm.getGameById(g1.id))
	}
}