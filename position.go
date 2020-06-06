package chess

// CastleRights is a string representation of rights to castle
type CastleRights string

func (cr CastleRights) isValid() bool {
	if cr == "-" {
		return true
	}
	if len(cr) < 1 || len(cr) > 4 {
		return false
	}
	validChars := map[rune]bool{'k': true, 'K': true, 'q': true, 'Q': true}
	for _, r := range cr {
		if _, ok := validChars[r]; !ok {
			return false
		}
	}
	return true
}

// Position represents a position in chess game
type Position struct {
	board         Board        // pieces on the board
	activeColor   Color        // White or Black
	epSquare      Square       // en-passant square (behind capturable pawn)
	halfMoveClock uint8        // a number of half moves since the last capture or pawn advance
	movesCount    uint8        // The number of the full move. It starts at 1, and is incremented after Black's move
	castleRight   CastleRights // rights of a castle
}

func (p *Position) Equal(position *Position) bool {
	if p == nil || position == nil {
		return false
	}
	return p.board.Equal(position.board) &&
		p.activeColor == position.activeColor &&
		p.epSquare == position.epSquare &&
		p.castleRight == position.castleRight &&
		p.movesCount == position.movesCount &&
		p.halfMoveClock == position.halfMoveClock
}

// NewStartPosition - initiates new start position
func NewStartPosition() *Position {
	board := map[Square]Piece{
		getSquare(FileA, Rank1): WhiteRook,
		getSquare(FileB, Rank1): WhiteKnight,
		getSquare(FileC, Rank1): WhiteBishop,
		getSquare(FileD, Rank1): WhiteQueen,
		getSquare(FileE, Rank1): WhiteKing,
		getSquare(FileF, Rank1): WhiteBishop,
		getSquare(FileG, Rank1): WhiteKnight,
		getSquare(FileH, Rank1): WhiteRook,
		getSquare(FileA, Rank2): WhitePawn,
		getSquare(FileB, Rank2): WhitePawn,
		getSquare(FileC, Rank2): WhitePawn,
		getSquare(FileD, Rank2): WhitePawn,
		getSquare(FileE, Rank2): WhitePawn,
		getSquare(FileF, Rank2): WhitePawn,
		getSquare(FileG, Rank2): WhitePawn,
		getSquare(FileH, Rank2): WhitePawn,
		getSquare(FileA, Rank8): BlackRook,
		getSquare(FileB, Rank8): BlackKnight,
		getSquare(FileC, Rank8): BlackBishop,
		getSquare(FileD, Rank8): BlackQueen,
		getSquare(FileE, Rank8): BlackKing,
		getSquare(FileF, Rank8): BlackBishop,
		getSquare(FileG, Rank8): BlackKnight,
		getSquare(FileH, Rank8): BlackRook,
		getSquare(FileA, Rank7): BlackPawn,
		getSquare(FileB, Rank7): BlackPawn,
		getSquare(FileC, Rank7): BlackPawn,
		getSquare(FileD, Rank7): BlackPawn,
		getSquare(FileE, Rank7): BlackPawn,
		getSquare(FileF, Rank7): BlackPawn,
		getSquare(FileG, Rank7): BlackPawn,
		getSquare(FileH, Rank7): BlackPawn,
	}
	return &Position{board: board, activeColor: White, epSquare: NoSquare, castleRight: "KQkq", movesCount: 1, halfMoveClock: 0}
}
