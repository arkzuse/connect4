package main

import (
	"github.com/google/uuid"
	"sync"
	"fmt"
)

type Game struct {
	id      uuid.UUID
	board   *Board
	players [2]uuid.UUID
	turn    int
	mu      sync.Mutex
}

func NewGame(player1 uuid.UUID, player2 uuid.UUID) *Game {
	return &Game{
		id:    uuid.New(),
		board: NewBoard(),
		players: [2]uuid.UUID{player1, player2},
		turn:  0,
	}
}

func (g *Game) playTurn(playerId uuid.UUID, col int) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if (g.players[g.turn] != playerId) {
		return fmt.Errorf("not player's turn")
	}

	err := g.board.placePiece(col, g.turn)

	g.turn = (g.turn + 1) % 2

	return err
}

func (g *Game) checkWinner() uuid.UUID {
	res := g.board.checkBoard()

	if res == -1 {
		return uuid.Nil
	}

	return g.players[res]
}
