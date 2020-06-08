package chess

import "testing"

func TestParseFEN(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	pos, err := ParseFEN(fen)
	if err != nil {
		t.Fatal(err)
	}
	if pos == nil {
		t.Fatal("position is null")
	}
	startPosition := NewStartPosition()
	if !startPosition.equal(pos) {
		t.Errorf("position parsed from FEN is not equal to expected position")
	}
}

func TestParseFEN2(t *testing.T) {
	fen := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"
	pos, err := ParseFEN(fen)
	if err != nil {
		t.Fatal(err)
	}
	if pos == nil {
		t.Fatal("position is null")
	}
	position := NewStartPosition()
	delete(position.board, E2)
	position.board[E4] = WhitePawn
	position.activeColor = Black
	position.epSquare = E3
	if !position.equal(pos) {
		t.Errorf("position parsed from FEN is not equal to expected position")
	}
}
