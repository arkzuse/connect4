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

func (g *Game) playTurn(player int, col int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.Board.placePiece(player, col)
}

func (g *Game) checkWinner() uuid.UUID {
	res := g.Board.checkBoard()

	if res == -1 {
		return uuid.UUID{}
	}

	return g.Players[res]
}
