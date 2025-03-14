package sudoku

import "slices"

type Cell struct {
	pos         int // 0〜80
	answer      int // 未確定は0
	candidates  []int
	isFixed     bool
	rowIndex    int
	columnIndex int
	squareIndex int
}

// func (c *Cell) Answer() int {
// 	if len(c.candidates) == 1 {
// 		return c.candidates[0]
// 	}
// 	return 0
// }

// func (c *Cell) IsFixed() bool {
// 	return c.Answer() != 0
// }

func NewCell(pos int, row int, column int, square int) *Cell {
	return &Cell{
		pos:         pos,
		answer:      0,
		candidates:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		isFixed:     false,
		rowIndex:    row,
		columnIndex: column,
		squareIndex: square,
	}
}

func (c *Cell) removeCandidate(candidate int) {
	if index := slices.Index(c.candidates, candidate); index > -1 {
		c.candidates = slices.Delete(c.candidates, index, index+1)
	}
}

func (c *Cell) fix(answer int) {
	c.answer = answer
	c.isFixed = true
}
