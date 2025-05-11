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

func (gm *GameManager) NewGame() *Game {
	g := NewGame()
	gm.games[g.id] = g
	return g
}

func (gm *GameManager) GetGame(id uuid.UUID) *Game {
	return gm.games[id]
}
