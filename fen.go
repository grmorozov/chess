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
			file += n - 1
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
	castleRights := CastleRights(segments[2])
	if !castleRights.isValid() {
		return nil, fmt.Errorf("castle rights is not in valid format: %s", castleRights)
	}
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
