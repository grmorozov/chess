package chess

// CastleRights is a string representation of rights to castle
type CastleRights int8

func parseCastleRights(s string) CastleRights {
	if s == "-" {
		return 0
	}
	c := 0
	for _, r := range s {
		switch r {
		case 'K': c |= 1
		case 'Q': c |= 2
		case 'k': c |= 4
		case 'q': c |= 8
		}
	}
	return CastleRights(c)
}

func (cr CastleRights) String() string {
	if cr == 0 {
		return "-"
	}
	s := ""
	if cr & 1 == 1 {
		s += "K"
	}
	if cr & 2 == 2 {
		s += "Q"
	}
	if cr & 4 == 4 {
		s += "k"
	}
	if cr & 8 == 8 {
		s += "q"
	}
	return s
}

func (cr CastleRights) update(m *Move, b Board) CastleRights {
	piece, ok := b[m.From]
	if !ok {
		return cr
	}
	if piece.Type() == King {
		if piece.Color() == White {
			return cr & 12
		}
		return cr & 3
	}
	if piece.Type() == Rook {
		switch m.From {
		case A1: return cr & 13
		case H1: return cr & 14
		case A8: return cr & 7
		case H8: return cr & 11
		}
	}
	return cr
}