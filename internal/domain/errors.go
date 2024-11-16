package domain

import "fmt"

type MarkAlreadyExistError struct {
	Row, Column int
}

func (e *MarkAlreadyExistError) Error() string {
	return fmt.Sprintf("position (%d, %d) already has a mark.", e.Row + 1, e.Column + 1)
}

type PositionOutOfBoundError struct {
	Row, Column int
}

func (e *PositionOutOfBoundError) Error() string {
	return fmt.Sprintf("position (%d, %d) is out of bounds.", e.Row +1 , e.Column + 1)
}
