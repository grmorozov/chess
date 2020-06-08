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