package chess

// CastleRights is a string representation of rights to castle
type CastleRights string

// Position represents a position in chess game
type Position struct {
	board         Board        // pieces on the board
	activeColor   Color        // White or Black
	epSquare      Square       // en-passant square (behind capturable pawn)
	halfMoveClock uint8        // a number of half moves since the last capture or pawn advance
	movesCount    uint8        // The number of the full move. It starts at 1, and is incremented after Black's move
	castleRight   CastleRights // rights of a castle
}
