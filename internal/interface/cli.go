package interfaces

import (
	"fmt"

	"github.com/LidoHon/tic-tac-toe/internal/domain"
	"github.com/LidoHon/tic-tac-toe/internal/usecase"
)

// ANSI color codes
const (
	Reset = "\033[0m"
	Blue  = "\033[34m" 
	Red   = "\033[31m" 
	BoardColor = "\033[37m" 
	Yellow = "\033[33m"
	Green = "\033[32m"
	cyan = "\033[36m"
	Purple = "\033[35m"
)

type CLI struct {
	Usecase *usecases.GameUsecase
}

// DrawBoard displays the game board with ANSI colors
func (cli *CLI) DrawBoard() {
	for i, row := range cli.Usecase.Game.Board {
		for j, square := range row {
			switch square {
			case domain.None:
				fmt.Print(BoardColor + "   " + Reset) 
			case domain.Cross:
				fmt.Print(Green + " X " + Reset)  
			case domain.Circle:
				fmt.Print(Yellow + " O " + Reset)   
			}

			if j < len(row)-1 {
				fmt.Print(Blue + " | " + Reset) 
			}
		}

		if i < len(cli.Usecase.Game.Board)-1 {
			fmt.Print("\n" + Blue + " ---+-----+--- " + Reset) 
		}
		fmt.Println()
	}
	fmt.Println()
}

// Start initializes the game loop
func (cli *CLI) Start() {
	for {
		cli.DrawBoard()
		var currentPlayer string
		if cli.Usecase.Game.PlayerTurn == domain.Cross {
			currentPlayer = Blue + "Player Cross" + Reset
		} else {
			currentPlayer = Purple + "Player Circle" + Reset
		}
		fmt.Printf("%s's turn. Enter row and column by separating with space (1-3): ", currentPlayer)
		var row, col int
		fmt.Scan(&row, &col)

		if err := cli.Usecase.PlaceMark(row-1, col-1); err != nil {
			fmt.Println(err)
			continue
		}
		result := cli.Usecase.CheckGameStatus()
		if result != domain.NoWinnerYet {
			cli.DrawBoard()
			switch result {
			case domain.CrossWon:
				fmt.Println(Green + "Cross wins!" + Reset)
			case domain.CircleWon:
				fmt.Println(Green + "Circle wins!" + Reset)
			case domain.Draw:
				fmt.Println(BoardColor + "It's a draw!" + Reset)
			}
			break
		}

		// Switch the turn to the next player
		cli.Usecase.SwitchTurn()
	}
}
