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
		{From: &Position{X: 4, Y: 0}, To: &Position{X: 7, Y: 0}},
	}

	game, err := setupGame(moves, t)
	if err != nil {
		t.Fatal(err)
	}

	if game.Board[7][6] == nil {
		t.Fatal("king position not set")
	}
	if game.Board[7][5] == nil {
		t.Fatal("rook position not set")
	}
	if game.Board[7][6].GetType() != King {
		t.Error("king not placed at the right position")
	}
	if game.Board[7][5].GetType() != Rook {
		t.Error("rook not placed at the right position")
	}

	if game.Board[0][6] == nil {
		t.Fatal("king position not set")
	}
	if game.Board[0][5] == nil {
		t.Fatal("rook position not set")
	}
	if game.Board[0][6].GetType() != King {
		t.Error("king not placed at the right position")
	}
	if game.Board[0][5].GetType() != Rook {
		t.Error("rook not placed at the right position")
	}

}

func TestPawnEnpassant(t *testing.T) {
	moves := []*Move{
		{From: &Position{X: 4, Y: 6}, To: &Position{X: 4, Y: 4}},
		{From: &Position{X: 2, Y: 1}, To: &Position{X: 2, Y: 3}},
		{From: &Position{X: 4, Y: 4}, To: &Position{X: 4, Y: 3}},
		{From: &Position{X: 5, Y: 1}, To: &Position{X: 5, Y: 3}},
		{From: &Position{X: 4, Y: 3}, To: &Position{X: 5, Y: 2}},
	}

	game, err := setupGame(moves, t)
	if err != nil {
		t.Fatal(err)
	}

	if game.Board[2][5] == nil {
		t.Fatal("pawn position not set")
	}

	if game.Board[2][5].GetType() != Pawn {
		t.Error("pawn not placed at the right position")
	}
}

func TestPawnCapture(t *testing.T) {
	moves := []*Move{
		{From: &Position{X: 4, Y: 6}, To: &Position{X: 4, Y: 4}},
		{From: &Position{X: 3, Y: 1}, To: &Position{X: 3, Y: 3}},
		{From: &Position{X: 4, Y: 4}, To: &Position{X: 3, Y: 3}},
	}

	game, err := setupGame(moves, t)
	if err != nil {
		t.Fatal(err)
	}

	if game.Board[3][3] == nil {
		t.Fatal("pawn position not set")
	}

	if game.Board[3][3].GetType() != Pawn && game.Board[3][3].GetColor() != White {
		t.Error("pawn not placed at the right position")
	}
}

func TestCheckmate(t *testing.T) {
	moves := []*Move{
		{From: &Position{X: 4, Y: 6}, To: &Position{X: 4, Y: 4}},
		{From: &Position{X: 4, Y: 1}, To: &Position{X: 4, Y: 3}},
		{From: &Position{X: 5, Y: 7}, To: &Position{X: 2, Y: 4}},
		{From: &Position{X: 1, Y: 0}, To: &Position{X: 2, Y: 2}},
		{From: &Position{X: 3, Y: 7}, To: &Position{X: 7, Y: 3}},
		{From: &Position{X: 6, Y: 0}, To: &Position{X: 5, Y: 2}},
		{From: &Position{X: 7, Y: 3}, To: &Position{X: 5, Y: 1}},
	}

	game, err := setupGame(moves, t)
	if err != nil {
		t.Fatal(err)
	}

	if game.Status != whiteCheckmate {
		t.Fatalf("game status is not white checkmate, Status: %s", game.Status)
	}
}
