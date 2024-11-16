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
	BoardColor = "\033[37m" // White for the board lines
	Yellow = "\033[33m"
	Green = "\033[32m"
	cyan = "\033[36m"
	Purple = "\033[35m"
)

// CLI represents the command-line interface for the game
type CLI struct {
	Usecase *usecases.GameUsecase
}

// DrawBoard displays the game board with ANSI colors
func (cli *CLI) DrawBoard() {
	// Loop through each row of the board
	for i, row := range cli.Usecase.Game.Board {
		// Loop through each column in the row
		for j, square := range row {
			// Print the square with appropriate colors
			switch square {
			case domain.None:
				fmt.Print(BoardColor + "   " + Reset) // Empty square
			case domain.Cross:
				fmt.Print(Green + " X " + Reset)  // Player X (Cross)
			case domain.Circle:
				fmt.Print(Yellow + " O " + Reset)   // Player O (Circle)
			}

			// Print a column separator unless it's the last column in the row
			if j < len(row)-1 {
				fmt.Print(Blue + " | " + Reset) // Column separator
			}
		}

		// Print a row separator unless it's the last row
		if i < len(cli.Usecase.Game.Board)-1 {
			fmt.Print("\n" + Blue + " ---+-----+--- " + Reset) // Row separator
		}

		// Move to the next line after each row
		fmt.Println()
	}

	// Add a blank line after the board for better readability
	fmt.Println()
}

// Start initializes the game loop
func (cli *CLI) Start() {
	for {
		cli.DrawBoard()

		// Display the current player with color
		var currentPlayer string
		if cli.Usecase.Game.PlayerTurn == domain.Cross {
			currentPlayer = Blue + "Player Cross" + Reset
		} else {
			currentPlayer = Purple + "Player Circle" + Reset
		}

		// Prompt the current player for their move
		fmt.Printf("%s's turn. Enter row and column by separating with space (1-3): ", currentPlayer)
		var row, col int
		fmt.Scan(&row, &col)

		// Place the mark and handle any errors
		if err := cli.Usecase.PlaceMark(row-1, col-1); err != nil {
			fmt.Println(err)
			continue
		}

		// Check for a winner or draw
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
