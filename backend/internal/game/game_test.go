package game

import (
	"testing"
)

func setupGame(moves []*Move, t *testing.T) (*Game, error) {
	g := NewGame()

	for _, move := range moves {
		t.Logf("handle move from %v to %v", move.From, move.To)
		err := g.HandleMove(move)
		if err != nil {
			return nil, err
		}
	}

	return g, nil
}

func TestKingCastling(t *testing.T) {
	moves := []*Move{ 
		{From: &Position{X: 4, Y: 6}, To: &Position{X: 4, Y: 4}},
		{From: &Position{X: 4, Y: 1}, To: &Position{X: 4, Y: 3}},
		{From: &Position{X: 6, Y: 7}, To: &Position{X: 5, Y: 5}},
		{From: &Position{X: 6, Y: 0}, To: &Position{X: 5, Y: 2}},
		{From: &Position{X: 5, Y: 7}, To: &Position{X: 2, Y: 4}},
		{From: &Position{X: 5, Y: 0}, To: &Position{X: 2, Y: 3}},
		{From: &Position{X: 4, Y: 7}, To: &Position{X: 7, Y: 7}},
	}

	game, err := setupGame(moves, t)
	if err != nil {
		t.Fatal(err)
	}

	if game.Board[7][6].GetType() != King {
		t.Error("king not placed at the right position")
	}
	if game.Board[7][5].GetType() != Rook {
		t.Error("rook not placed at the right position")
	}
}

