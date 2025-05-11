package main

import (
	"github.com/google/uuid"
)

type Player struct {
	id      uuid.UUID
	name	string
}

func NewPlayer(name string) *Player {
	return &Player{
		id: uuid.New(),
		name: name,
	}
}
