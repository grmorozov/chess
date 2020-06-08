package chess

// Move represents a move in chess game
type Move struct {
	From      Square
	To        Square
	Promotion PieceType
}
