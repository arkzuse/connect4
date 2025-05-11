package main

import (
	"testing"
)

func TestGameManager(t * testing.T) {
	gm := NewGameManager()

	g1 := gm.newGame()

	// get game
	if gm.getGameById(g1.id) != g1 {
		t.Errorf("Expected game %v, got %v", g1, gm.games[g1.id])
	}

	// remove game
	gm.removeGame(g1.id)
	if gm.getGameById(g1.id) != nil {
		t.Errorf("Expecter player to be nil, got %v", gm.getGameById(g1.id))
	}
}