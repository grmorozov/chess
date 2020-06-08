package chess

type PieceType int8
// Piece
const (
	NoPieceType PieceType = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p PieceType) Name() string {
	return [7]string{"No piece", "Pawn", "Knight", "Bishop", "Rook", "Queen", "King"}[int(p)]
}

// Piece represent chess piece
type Piece int8

// Color - retrieves piece color
func (p Piece) Color() Color { return Color(p >> 3) }

// Type - retrieves piece type
func (p Piece) Type() PieceType { return PieceType(int8(p) & 0b111) }

func (p Piece) Name() string {
	return p.Color().Name() + p.Type().Name()
}

func getPiece(color Color, pieceType PieceType) Piece {
	return Piece(int8(color) << 3 + int8(pieceType))
}

// Pieces
const (
	NoPiece = Piece(int8(NoColor) << 3 + int8(NoPieceType))
	WhitePawn = Piece(int8(White) << 3 + int8(Pawn))
	WhiteKnight = Piece(int8(White) << 3 + int8(Knight))
	WhiteBishop = Piece(int8(White) << 3 + int8(Bishop))
	WhiteRook = Piece(int8(White) << 3 + int8(Rook))
	WhiteQueen = Piece(int8(White) << 3 + int8(Queen))
	WhiteKing = Piece(int8(White) << 3 + int8(King))
	BlackPawn = Piece(int8(Black) << 3 + int8(Pawn))
	BlackKnight = Piece(int8(Black) << 3 + int8(Knight))
	BlackBishop = Piece(int8(Black) << 3 + int8(Bishop))
	BlackRook = Piece(int8(Black) << 3 + int8(Rook))
	BlackQueen = Piece(int8(Black) << 3 + int8(Queen))
	BlackKing = Piece(int8(Black) << 3 + int8(King))
)