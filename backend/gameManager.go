package main

import "github.com/google/uuid"

type GameManager struct {
	games map[uuid.UUID]*Game
}

func NewGameManager() *GameManager {
	return &GameManager{
		games: make(map[uuid.UUID]*Game),
	}
}

func (gm *GameManager) newGame(player1 uuid.UUID, player2 uuid.UUID) *Game {
	g := NewGame(player1, player2)
	gm.games[g.id] = g
	return g
}

func (gm *GameManager) getGameById(id uuid.UUID) *Game {
	return gm.games[id]
}

func (gm *GameManager) removeGame(id uuid.UUID) {
	delete(gm.games, id)
}
