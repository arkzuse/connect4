package main

import (
	"testing"
)

func TestPlayerManager(t *testing.T) {
	pm := NewPlayerManager()

	p1 := pm.newPlayer("Alice")

	if (pm.getPlayerById(p1.id) != p1) {
		t.Errorf("Expected player %v, got %v", p1, pm.players[p1.id])
	}

	// remove player
	pm.removePlayer(p1.id)
	if pm.getPlayerById(p1.id) != nil {
		t.Errorf("Expected player to be nil, got %v", pm.getPlayerById(p1.id))
	}
}
