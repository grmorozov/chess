package chess

// Board represents chess board with pieces on it
type Board map[Square]Piece

func (b Board) equal(board Board) bool {
	if len(b) != len(board) {
		return false
	}
	for s, p := range board {
		if piece, ok := b[s]; !ok || piece != p {
			return false
		}
	}
	return true
}

// Copy creates a copy of a board
func (b Board) copy() Board {
	board := Board{}
	for s, p := range b {
		board[s] = p
	}
	return board
}

// Draw returns visual representation of the board to display in console
func (b Board) Draw() string {
	s := "\n \u2001A\u2001B\u2001C\u2001D\u2001E\u2001F\u2001G\u2001H\n"
	for r := 7; r >= 0; r-- {
		s += Rank(r).String() + "\u2001"
		for f := 0; f < numOfSquaresInRow; f++ {
			p := b[getSquare(File(f), Rank(r))]
			if p == NoPiece {
				s += "\u2001"
			} else {
				s += p.String()
			}
			s += " "
		}
		s += "\n"
	}
	return s
}
