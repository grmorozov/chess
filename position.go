package chess



// Position represents a position in chess game
type Position struct {
	board         Board        // pieces on the board
	activeColor   Color        // White or Black
	epSquare      Square       // en-passant square (behind capturable pawn)
	halfMoveClock uint8        // a number of half moves since the last capture or pawn advance
	movesCount    uint8        // The number of the full move. It starts at 1, and is incremented after Black's move
	castleRight   CastleRights // rights of a castle
}

func (p *Position) equal(position *Position) bool {
	if p == nil || position == nil {
		return false
	}
	return p.board.equal(position.board) &&
		p.activeColor == position.activeColor &&
		p.epSquare == position.epSquare &&
		p.castleRight == position.castleRight &&
		p.movesCount == position.movesCount &&
		p.halfMoveClock == position.halfMoveClock
}

// Copy creates a new copy of the position
func (p *Position) copy() *Position {
	return &Position{board: p.board.copy(), activeColor: p.activeColor, epSquare: p.epSquare,
		castleRight: p.castleRight, movesCount: p.movesCount, halfMoveClock: p.halfMoveClock}
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
	return &Position{board: board, activeColor: White, epSquare: NoSquare, castleRight: parseCastleRights("KQkq"), movesCount: 1, halfMoveClock: 0}
}

// MakeMove - apply move to the position and returns new position, no validation
func (p *Position) MakeMove(m *Move) *Position {
	board := p.board.copy()
	piece := board[m.From]
	isCapture := false
	if targetPiece, ok := board[m.To]; ok && targetPiece.Type() != NoPieceType {
		isCapture = true
	}
	board[m.To] = board[m.From]
	delete(board, m.From)

	if m.Promotion != NoPieceType {
		board[m.To] = getPiece(piece.Color(), m.Promotion)
	}

	// if castle, move rook (classic chess only, doesn't work for chess960)
	if piece.Type() == King && (int(m.From.File())+int(m.To.File()))%2==0 {
		var from, to Square
		if m.From.File() < m.To.File() {	// king side
			from = getSquare(FileH, m.To.Rank())
			to = getSquare(FileF, m.To.Rank())
		} else {	// queen side
			from = getSquare(FileA, m.To.Rank())
			to = getSquare(FileD, m.To.Rank())
		}
		// move rook
		board[to] = board[from]
		delete(board, from)
	}

	// remove en-passant pawn
	if piece.Type() == Pawn && m.To == p.epSquare {
		delete(board, getSquare(p.epSquare.File(), m.From.Rank()))
	}

	epSquare := NoSquare
	if piece.Type() == Pawn && (int(m.From.Rank())+int(m.To.Rank()))%2 == 0 {
		rank := Rank((int(m.From.Rank()) + int(m.To.Rank())) / 2)
		epSquare = getSquare(m.From.File(), rank)
	}

	movesCount := p.movesCount
	if p.activeColor == Black {
		movesCount++
	}

	halfMove := p.halfMoveClock + 1
	if piece.Type() == Pawn || isCapture {
		halfMove = 0
	}

	return &Position{
		board:         board,
		activeColor:   p.activeColor.Other(),
		epSquare:      epSquare,
		halfMoveClock: halfMove,
		movesCount:    movesCount,
		castleRight:   p.castleRight.update(m, p.board),
	}
}