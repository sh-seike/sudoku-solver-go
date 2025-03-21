package sudoku

import "slices"

// Cellは数独の１マスを表します。
// セルの位置、確定された場合の答え、
// 答えが確定されていない場合の候補数字のリスト、
// およびセルが属する行、列、スクエアのインデックスを含みます。
type Cell struct {
	pos         int // 0〜80
	answer      int // 未確定は0
	candidates  []int
	isFixed     bool
	rowIndex    int
	columnIndex int
	squareIndex int
}

// answerやisFixedを以下のようなメソッドで定義することもできるが、
// これらは何度も繰り返し呼び出されるため、パフォーマンスを考慮してフィールドとして定義する。
// 状態を正しく更新しないとバグの原因になるため注意。
// func (c *Cell) Answer() int {
// 	if len(c.candidates) == 1 {
// 		return c.candidates[0]
// 	}
// 	return 0
// }

// func (c *Cell) IsFixed() bool {
// 	return c.Answer() != 0
// }

// NewCellは指定された位置、行、列、およびスクエアインデックスで新しいCellインスタンスを作成します。
// このメソッドは、セルをすべての可能な候補（1〜9）で初期化し、answerを0に設定し、答えが確定されていないとマークします。
//
// パラメータ:
//   - pos: Sudokuグリッド内のセルの位置
//   - row: セルの行インデックス
//   - column: セルの列インデックス
//   - square: セルのスクエアインデックス
//
// 戻り値:
//   - 新しく作成されたCellインスタンスへのポインタ
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

// removeCandidateは、セルの候補リストから指定された候補を削除します。
//
// パラメータ:
//   - candidate: 削除する候補
func (c *Cell) removeCandidate(candidate int) {
	if index := slices.Index(c.candidates, candidate); index > -1 {
		c.candidates = slices.Delete(c.candidates, index, index+1)
	}
}

// fixはセルの答えを指定された値に設定し、セルを確定済みとしてマークします。
// パラメータ:
//   - answer: セルの答えとして設定する値
func (c *Cell) fix(answer int) {
	c.answer = answer
	c.isFixed = true
}
