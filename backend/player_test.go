package main

import (
	"testing"
	"github.com/google/uuid"
)

func TestPlayer(t *testing.T) {
	p := NewPlayer("Alice")

	if p.name != "Alice" {
		t.Errorf("Expected player name to be 'Alice', got '%s'", p.name)
	}

	if p.id == uuid.Nil {
		t.Errorf("Expected player ID to be non-nil, got nil")
	}
}
