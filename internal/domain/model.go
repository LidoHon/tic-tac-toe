package domain

type Player int
type GameResult int

// Player constants
const (
	None   Player = iota // this means like 0 and cross get the next value means 1 and circle get the next value means 2
	Cross                
	Circle               
)

// GameResult constants
const (
	NoWinnerYet GameResult = iota
	CrossWon
	CircleWon
	Draw
)

// GameState represents the current state of the game

// Next method determines the next player's turn
func (p Player) Next() Player {
	if p == Cross {
		return Circle
	}
	return Cross
}
