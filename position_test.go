package chess

import "testing"

func TestPosition_MakeMove(t *testing.T) {
	type testCase struct {
		move *Move
		fen  string
	}
	position := NewStartPosition()
	testCases := []testCase{
		{&Move{E2, E4, NoPieceType}, "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1"},
		{&Move{C7, C5, NoPieceType}, "rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2"},
		{&Move{G1, F3, NoPieceType}, "rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2"},
	}
	for _, tc := range testCases {
		position = position.MakeMove(tc.move)
		if expected, err := ParseFEN(tc.fen); err != nil {
			t.Errorf("fail to parse fen %s: %v", tc.fen, err)
		} else if !position.equal(expected) {
			t.Errorf("position is not equal expected")
		}
	}
}
