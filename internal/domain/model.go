package domain

type Player int
type GameResult int

const (
	None   Player = iota // this means like 0 and cross get the next value means 1 and circle get the next value means 2
	Cross                
	Circle               
)

const (
	NoWinnerYet GameResult = iota
	CrossWon
	CircleWon
	Draw
)

func (p Player) Next() Player {
	if p == Cross {
		return Circle
	}
	return Cross
}
