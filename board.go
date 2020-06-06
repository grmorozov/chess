package chess

// Board represents chess board with pieces on it
type Board map[Square]Piece

func (b Board) Equal(board Board) bool {
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
