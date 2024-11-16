package main

import (
	"github.com/LidoHon/tic-tac-toe/internal/domain"
	interfaces "github.com/LidoHon/tic-tac-toe/internal/interface"
	usecases "github.com/LidoHon/tic-tac-toe/internal/usecase"
)


func main() {
	gameState := &domain.GameState{PlayerTurn: domain.Cross}
	gameUsecase := &usecases.GameUsecase{Game: gameState}
	cli := &interfaces.CLI{Usecase: gameUsecase}

	cli.Start()
}
