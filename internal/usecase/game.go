package usecases

import (
	"github.com/LidoHon/tic-tac-toe/internal/domain"

)

type GameUsecase struct {
	Game *domain.GameState
}

func (u *GameUsecase) PlaceMark(row, col int) error {
	return u.Game.PlaceMark(row, col)
}


func (u *GameUsecase) CheckGameStatus() domain.GameResult {
	return u.Game.CheckForWinner()
}

func (u *GameUsecase) SwitchTurn() {
	u.Game.TurnPlayer = u.Game.TurnPlayer.Next()
}
