package chess

import "testing"

func TestPiece_Color(t *testing.T) {
	testCases := map[Piece]Color{
		WhiteKing:   White,
		BlackQueen:  Black,
		NoPiece:     NoColor,
		WhitePawn:   White,
		BlackKnight: Black,
	}
	for piece, color := range testCases {
		actualColor := piece.Color()
		if actualColor != color {
			t.Errorf("expected color for %v is %s actual %s", piece, color, actualColor)
		}
	}
}

func TestPiece_Type(t *testing.T) {
	testCases := map[Piece]PieceType{
		WhiteKing:   King,
		BlackQueen:  Queen,
		NoPiece:     NoPieceType,
		WhitePawn:   Pawn,
		BlackKnight: Knight,
	}
	for piece, pieceType := range testCases {
		actualType := piece.Type()
		if actualType != pieceType {
			t.Errorf("expected type for %v is %s actual %s", piece, pieceType.Name(), actualType.Name())
		}
	}
}
