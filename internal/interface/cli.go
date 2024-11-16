package interfaces

import (
	"fmt"
	"github.com/LidoHon/tic-tac-toe/internal/domain"
	usecases "github.com/LidoHon/tic-tac-toe/internal/usecase"
)

// ANSI color codes
const (
	ResetColor = "\033[0m"
	BlueColor  = "\033[34m" // Blue for Cross
	RedColor   = "\033[31m" // Red for Circle
	BoardColor = "\033[37m" // White for the board lines
)

// CLI represents the command-line interface for the game
type CLI struct {
	Usecase *usecases.GameUsecase
}

// DrawBoard displays the game board with ANSI colors
// DrawBoard displays the game board with ANSI colors
func (cli *CLI) DrawBoard() {
	// Loop through each row of the board
	for i, row := range cli.Usecase.Game.Board {
		// Loop through each column in the row
		for j, square := range row {
			// Print the square with appropriate colors
			switch square {
			case domain.None:
				fmt.Print(BoardColor + "   " + ResetColor) // Empty square
			case domain.Cross:
				fmt.Print(BlueColor + "X " + ResetColor)  // Player X (Cross)
			case domain.Circle:
				fmt.Print(RedColor + "O " + ResetColor)   // Player O (Circle)
			}

			// Print a column separator unless it's the last column in the row
			if j < len(row)-1 {
				fmt.Print(BoardColor + " | " + ResetColor) // Column separator
			}
		}

		// Print a row separator unless it's the last row
		if i < len(cli.Usecase.Game.Board)-1 {
			fmt.Print("\n" + BoardColor + " ---+-----+--- " + ResetColor) // Row separator
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
		if cli.Usecase.Game.TurnPlayer == domain.Cross {
			currentPlayer = BlueColor + "Player Cross" + ResetColor
		} else {
			currentPlayer = RedColor + "Player Circle" + ResetColor
		}

		// Prompt the current player for their move
		fmt.Printf("%s's turn. Enter row and column (1-3): ", currentPlayer)
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
				fmt.Println(BlueColor + "Cross wins!" + ResetColor)
			case domain.CircleWon:
				fmt.Println(RedColor + "Circle wins!" + ResetColor)
			case domain.Draw:
				fmt.Println(BoardColor + "It's a draw!" + ResetColor)
			}
			break
		}

		// Switch the turn to the next player
		cli.Usecase.SwitchTurn()
	}
}
