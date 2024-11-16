package domain

// PlaceMark places a mark for the current player on the board
func (state *GameState) PlaceMark(row, col int) error {
	if row < 0 || col < 0 || row >= len(state.Board) || col >= len(state.Board[row]) {
		return &PositionOutOfBoundError{Row: row, Column: col}
	}
	if state.Board[row][col] != None {
		return &MarkAlreadyExistError{Row: row, Column: col}
	}

	state.Board[row][col] = state.TurnPlayer // Use Player type directly
	return nil
}

// CheckForWinner checks the current board state for a winner
func (state *GameState) CheckForWinner() GameResult {
	// Lambda function to check a line on the board
	checkLine := func(r, c, dr, dc int) GameResult {
		first := state.Board[r][c]
		if first == None {
			return NoWinnerYet
		}
		for i := 1; i < 3; i++ {
			r, c = r+dr, c+dc
			if state.Board[r][c] != first {
				return NoWinnerYet
			}
		}
		if first == Cross {
			return CrossWon
		}
		return CircleWon
	}

	// Check rows and columns
	for i := 0; i < 3; i++ {
		if result := checkLine(i, 0, 0, 1); result != NoWinnerYet {
			return result
		}
		if result := checkLine(0, i, 1, 0); result != NoWinnerYet {
			return result
		}
	}
	// Check diagonals
	if result := checkLine(0, 0, 1, 1); result != NoWinnerYet {
		return result
	}
	if result := checkLine(0, 2, 1, -1); result != NoWinnerYet {
		return result
	}

	// Check for draw
	for _, row := range state.Board {
		for _, square := range row {
			if square == None {
				return NoWinnerYet
			}
		}
	}
	return Draw
}
