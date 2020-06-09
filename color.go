package chess

// Color represents the color of a chess piece or square.
type Color int8

const (
	// NoColor represents no color
	NoColor Color = iota
	// White represents the color white
	White
	// Black represents the color black
	Black
)

func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	}
	return "-"
}

func (c Color) Name() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	return "No Color"
}

// Other returns the opposite color of the receiver.
func (c Color) Other() Color {
	if c == White {
		return Black
	} else if c == Black {
		return White
	}
	return NoColor
}
