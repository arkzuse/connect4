package main

import "github.com/google/uuid"

type PlayerManager struct {
	players	map[uuid.UUID]*Player
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		players: make(map[uuid.UUID]*Player),
	}
}

func (pm *PlayerManager) newPlayer(name string) *Player {
	p := NewPlayer(name)
	pm.players[p.id] = p
	return p
}

func (pm *PlayerManager) getPlayerById(id uuid.UUID) *Player {
	if player, ok := pm.players[id]; ok {
		return player
	}

	return nil
}

func (pm *PlayerManager) removePlayer(id uuid.UUID) {
	delete(pm.players, id)
}
