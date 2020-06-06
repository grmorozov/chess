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
	if !startPosition.Equal(pos) {
		t.Errorf("position parsed from FEN is not equal to expected position")
	}
}
