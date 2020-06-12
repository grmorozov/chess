package chess

import "fmt"

type gameState struct {
	position   *Position
	nextMove   *Move
	nextState  *gameState
	prevState  *gameState
	variants   map[Move]*gameState
	textBefore string
	textAfter  string
}

// Game represents a full chess game
type Game struct {
	state *gameState
	tags  map[string]string
}

// Move adds next main line move to the game and moves game to new state
func (g *Game) Move(move *Move) error {
	if g.state.nextMove != nil {
		return fmt.Errorf("another move was already made in current position")
	}
	if !g.state.position.IsValidMove(move) {
		return fmt.Errorf("move is not valid in current position")
	}
	g.state.nextMove = move
	g.state.nextState = &gameState{
		position:  g.state.position.MakeMove(move),
		prevState: g.state,
	}
	g.state = g.state.nextState
	return nil
}

// AddVariant adds an alternative move and moves game to the new state
func (g *Game) AddVariant(move *Move) error {
	if g.state.variants == nil {
		g.state.variants = make(map[Move]*gameState, 1)
	}
	if state, ok := g.state.variants[*move]; ok {
		g.state = state
		return nil
	}
	if !g.state.position.IsValidMove(move) {
		return fmt.Errorf("move is not valid in current position")
	}
	state := &gameState{
		position:  g.state.position.MakeMove(move),
		prevState: g.state,
	}
	g.state.variants[*move] = state
	g.state = state
	return nil
}

// NextState moves game state to next position along the main line; if game is in final position, does nothing
func (g *Game) NextState() *Game {
	if g.state.nextState != nil {
		g.state = g.state.nextState
	}
	return g
}

// PrevState moves game to the previous position; if game is in initial position, does nothing
func (g *Game) PrevState() *Game {
	if g.state.prevState != nil {
		g.state = g.state.prevState
	}
	return g
}
