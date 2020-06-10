package chess

import (
	"fmt"
	"strconv"
	"strings"
)

var charToPieceMap = map[rune]Piece{
	'p': BlackPawn,
	'P': WhitePawn,
	'q': BlackQueen,
	'Q': WhiteQueen,
	'k': BlackKing,
	'K': WhiteKing,
	'r': BlackRook,
	'R': WhiteRook,
	'b': BlackBishop,
	'B': WhiteBishop,
	'n': BlackKnight,
	'N': WhiteKnight,
}

var pieceToCharMap = map[Piece]rune{
	BlackPawn:   'p',
	WhitePawn:   'P',
	BlackQueen:  'q',
	WhiteQueen:  'Q',
	BlackKing:   'k',
	WhiteKing:   'K',
	BlackRook:   'r',
	WhiteRook:   'R',
	BlackBishop: 'b',
	WhiteBishop: 'B',
	BlackKnight: 'n',
	WhiteKnight: 'N',
}

var charToDigitMap = map[rune]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
}

var digits = "012345678"

func parseBoard(s string) (Board, error) {
	pieces := make(map[Square]Piece)
	rank, file := 7, 0
	for _, r := range s {
		if r == '/' {
			rank--
			file = 0
		} else if piece, ok := charToPieceMap[r]; ok {
			square := getSquare(File(file), Rank(rank))
			pieces[square] = piece
			file++
		} else if n, ok := charToDigitMap[r]; ok {
			file += n
		} else {
			return nil, fmt.Errorf("failed to parse symbol: %c", r)
		}
	}
	return pieces, nil
}

// ParseFEN - parses Position from FEN format (https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation)
func ParseFEN(fen string) (*Position, error) {
	segments := strings.Split(fen, " ")
	if len(segments) != 6 {
		return nil, fmt.Errorf("expected 6 segemnts, got %d", len(segments))
	}
	board, err := parseBoard(segments[0])
	if err != nil {
		return nil, err
	}
	castleRights := parseCastleRights(segments[2])
	epSquare, ok := strToSquareMap[segments[3]]
	if !ok {
		return nil, fmt.Errorf("fail to parse '%s' as an en-passant square", segments[3])
	}
	position := &Position{board: board, activeColor: White, epSquare: epSquare, castleRight: castleRights}
	if segments[1] == "b" {
		position.activeColor = Black
	}

	halfMoveClock, err := strconv.Atoi(segments[4])
	if err != nil {
		return nil, fmt.Errorf("fail to parse halfmoves: %v", err)
	}
	position.halfMoveClock = uint8(halfMoveClock)
	fullMovesNumber, err := strconv.Atoi(segments[5])
	if err != nil {
		return nil, fmt.Errorf("fail to parse full moves number: %v", err)
	}
	position.movesCount = uint8(fullMovesNumber)

	return position, nil
}

func GenerateFEN(position *Position) string {
	s := ""
	// board
	for r := 7; r >= 0; r-- {
		n := 0
		for f := 0; f < numOfSquaresInRow; f++ {
			piece := position.board[getSquare(File(f), Rank(r))]
			if piece == NoPiece {
				n++
			} else {
				if n > 0 {
					s += string(digits[n])
				}
				s += string(pieceToCharMap[piece])
				n = 0
			}
		}
		if n > 0 {
			s += string(digits[n])
		}
		if r > 0 {
			s += "/"
		}
	}

	epSquare := "-"
	if position.epSquare != NoSquare {
		epSquare = position.epSquare.String()
	}
	return fmt.Sprintf("%s %s %s %s %d %d", s, position.activeColor, position.castleRight, epSquare,
		position.halfMoveClock, position.movesCount)
}
