package domain

type Player int
type GameResult int

// Player constants
const (
	None   Player = iota // Represents an empty square or no player
	Cross                // Represents the Cross player
	Circle               // Represents the Circle player
)

// GameResult constants
const (
	NoWinnerYet GameResult = iota
	CrossWon
	CircleWon
	Draw
)

// GameState represents the current state of the game
type GameState struct {
	Board      [3][3]Player // Each square on the board stores the current player's mark
	TurnPlayer Player       // Tracks whose turn it is
}

// Next method determines the next player's turn
func (p Player) Next() Player {
	if p == Cross {
		return Circle
	}
	return Cross
}
