package main

import (
	"github.com/google/uuid"
	"sync"
)

type Game struct {
	ID      uuid.UUID
	Board   *Board
	Players [2]uuid.UUID
	turn    int
	mu      sync.Mutex
}

func NewGame() *Game {
	return &Game{
		ID:    uuid.New(),
		Board: newBoard(),
		turn:  0,
	}
}
