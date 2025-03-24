package sudoku

import (
	"fmt"
	"slices"
	"strings"
)

// var board [][]int

// type Cell struct {
// 	candidates  []int
// }

// type Board struct {
// 	grid     []Cell
// }
//
// インデックスを9で割った商が行、余りが列になる
//   ||  0 |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |
// -------------------------------------------------
// 0 ||  0 |  1 |  2 |  3 |  4 |  5 |  6 |  7 |  8 |
// 1 ||  9 | 10 | 11 | 12 | 13 | 14 | 15 | 16 | 17 |
// 2 || 18 | 19 | 20 | 21 | 22 | 23 | 24 | 25 | 26 |
// 3 || 27 | 28 | 29 | 30 | 31 | 32 | 33 | 34 | 35 |
// 4 || 36 | 37 | 38 | 39 | 40 | 41 | 42 | 43 | 44 |
// 5 || 45 | 46 | 47 | 48 | 49 | 50 | 51 | 52 | 53 |
// 6 || 54 | 55 | 56 | 57 | 58 | 59 | 60 | 61 | 62 |
// 7 || 63 | 64 | 65 | 66 | 67 | 68 | 69 | 70 | 71 |
// 8 || 72 | 73 | 74 | 75 | 76 | 77 | 78 | 79 | 80 |

// Boardは数独の盤面を表します。
// セルのグリッド、行、列、3x3のマス目を含みます。
// debugフィールドはデバッグ情報の有効/無効を制御します。
//
// フィールド:
//   - grid: 数独の全てのマスを表すCellのslice
//   - rows: 盤面の行を表すCellへのポインタの2次元slice
//   - columns: 盤面の列を表すCellへのポインタの2次元slice
//   - squares: 盤面の3x3のマス目を表すCellへのポインタの2次元slice
//   - debug: デバッグ情報の有効/無効を制御するbool値
type Board struct {
	grid     []Cell
	rows     [][]*Cell
	collumns [][]*Cell
	squares  [][]*Cell
	debug    bool
}

// NewBoardは81個のCellで初期化された数独の盤面を作成します。
// 各Cellにはそれぞれの行、列、および3x3のマス目のインデックスが割り当てられます。
// この関数は新しく作成されたBoardのポインタを返します。
//
// パラメータ:
//   - debug: デバッグモードを有効にするかどうかを示すフラグ
//
// 戻り値:
//   - *Board: 新しく作成されたBoardインスタンスへのポインタ
func NewBoard(debug bool) *Board {
	grid := make([]Cell, 81)

	for i := range grid {
		rowIndex := i / 9
		columnIndex := i % 9
		squereIndex := (rowIndex/3)*3 + columnIndex/3
		grid[i] = *NewCell(i, rowIndex, columnIndex, squereIndex)
	}

	return NewBoardWithGrid(grid, debug)
}

// NewBoardWithGridは提供されたCellのグリッドとデバッグフラグを使用して新しいBoardを作成します。
//
// パラメータ:
//   - g: 初期グリッドを表すCellのslice
//   - d: デバッグモードを有効にするかどうかを示すフラグ
//
// 戻り値:
//   - 新しく作成されたBoardのポインタ
func NewBoardWithGrid(g []Cell, d bool) *Board {
	grid := make([]Cell, 81)
	copy(grid, g)
	for i := range g {
		grid[i].candidates = make([]int, len(g[i].candidates))
		copy(grid[i].candidates, g[i].candidates)
	}
	rows := make([][]*Cell, 9)
	collumns := make([][]*Cell, 9)
	squares := make([][]*Cell, 9)
	rows[0] = []*Cell{&grid[0], &grid[1], &grid[2], &grid[3], &grid[4], &grid[5], &grid[6], &grid[7], &grid[8]}
	rows[1] = []*Cell{&grid[9], &grid[10], &grid[11], &grid[12], &grid[13], &grid[14], &grid[15], &grid[16], &grid[17]}
	rows[2] = []*Cell{&grid[18], &grid[19], &grid[20], &grid[21], &grid[22], &grid[23], &grid[24], &grid[25], &grid[26]}
	rows[3] = []*Cell{&grid[27], &grid[28], &grid[29], &grid[30], &grid[31], &grid[32], &grid[33], &grid[34], &grid[35]}
	rows[4] = []*Cell{&grid[36], &grid[37], &grid[38], &grid[39], &grid[40], &grid[41], &grid[42], &grid[43], &grid[44]}
	rows[5] = []*Cell{&grid[45], &grid[46], &grid[47], &grid[48], &grid[49], &grid[50], &grid[51], &grid[52], &grid[53]}
	rows[6] = []*Cell{&grid[54], &grid[55], &grid[56], &grid[57], &grid[58], &grid[59], &grid[60], &grid[61], &grid[62]}
	rows[7] = []*Cell{&grid[63], &grid[64], &grid[65], &grid[66], &grid[67], &grid[68], &grid[69], &grid[70], &grid[71]}
	rows[8] = []*Cell{&grid[72], &grid[73], &grid[74], &grid[75], &grid[76], &grid[77], &grid[78], &grid[79], &grid[80]}

	collumns[0] = []*Cell{&grid[0], &grid[9], &grid[18], &grid[27], &grid[36], &grid[45], &grid[54], &grid[63], &grid[72]}
	collumns[1] = []*Cell{&grid[1], &grid[10], &grid[19], &grid[28], &grid[37], &grid[46], &grid[55], &grid[64], &grid[73]}
	collumns[2] = []*Cell{&grid[2], &grid[11], &grid[20], &grid[29], &grid[38], &grid[47], &grid[56], &grid[65], &grid[74]}
	collumns[3] = []*Cell{&grid[3], &grid[12], &grid[21], &grid[30], &grid[39], &grid[48], &grid[57], &grid[66], &grid[75]}
	collumns[4] = []*Cell{&grid[4], &grid[13], &grid[22], &grid[31], &grid[40], &grid[49], &grid[58], &grid[67], &grid[76]}
	collumns[5] = []*Cell{&grid[5], &grid[14], &grid[23], &grid[32], &grid[41], &grid[50], &grid[59], &grid[68], &grid[77]}
	collumns[6] = []*Cell{&grid[6], &grid[15], &grid[24], &grid[33], &grid[42], &grid[51], &grid[60], &grid[69], &grid[78]}
	collumns[7] = []*Cell{&grid[7], &grid[16], &grid[25], &grid[34], &grid[43], &grid[52], &grid[61], &grid[70], &grid[79]}
	collumns[8] = []*Cell{&grid[8], &grid[17], &grid[26], &grid[35], &grid[44], &grid[53], &grid[62], &grid[71], &grid[80]}

	squares[0] = []*Cell{&grid[0], &grid[1], &grid[2], &grid[9], &grid[10], &grid[11], &grid[18], &grid[19], &grid[20]}
	squares[1] = []*Cell{&grid[3], &grid[4], &grid[5], &grid[12], &grid[13], &grid[14], &grid[21], &grid[22], &grid[23]}
	squares[2] = []*Cell{&grid[6], &grid[7], &grid[8], &grid[15], &grid[16], &grid[17], &grid[24], &grid[25], &grid[26]}
	squares[3] = []*Cell{&grid[27], &grid[28], &grid[29], &grid[36], &grid[37], &grid[38], &grid[45], &grid[46], &grid[47]}
	squares[4] = []*Cell{&grid[30], &grid[31], &grid[32], &grid[39], &grid[40], &grid[41], &grid[48], &grid[49], &grid[50]}
	squares[5] = []*Cell{&grid[33], &grid[34], &grid[35], &grid[42], &grid[43], &grid[44], &grid[51], &grid[52], &grid[53]}
	squares[6] = []*Cell{&grid[54], &grid[55], &grid[56], &grid[63], &grid[64], &grid[65], &grid[72], &grid[73], &grid[74]}
	squares[7] = []*Cell{&grid[57], &grid[58], &grid[59], &grid[66], &grid[67], &grid[68], &grid[75], &grid[76], &grid[77]}
	squares[8] = []*Cell{&grid[60], &grid[61], &grid[62], &grid[69], &grid[70], &grid[71], &grid[78], &grid[79], &grid[80]}

	return &Board{
		grid:     grid,
		rows:     rows,
		collumns: collumns,
		squares:  squares,
		debug:    d,
	}
}

// Updateは指定された位置に解答を設定し、Cellを確定としてマークします。
// 同時に同じ行、列、および3x3のマス目の固定されていない全てのCellから解答を候補から削除します。
//
// パラメータ:
//   - pos: 盤面の更新する位置
//   - answer: 指定された位置に設定する値
func (b *Board) Update(pos int, answer int) {
	cell := &b.grid[pos]
	cell.fix(answer)

	row := b.rows[cell.rowIndex]
	for _, c := range row {
		if !c.isFixed {
			c.removeCandidate(answer)
		}
	}
	collumn := b.collumns[cell.columnIndex]
	for _, c := range collumn {
		if !c.isFixed {
			c.removeCandidate(answer)
		}
	}
	square := b.squares[cell.squareIndex]
	for _, c := range square {
		if !c.isFixed {
			c.removeCandidate(answer)
		}
	}
}

// Printは現在の数独の盤面を出力します。
// 盤面がデバッグモードでなく、forceパラメータがfalseの場合、この関数は出力せずに終了します。
//
// パラメータ:
//   - force - 盤面がデバッグモードでなくても強制的に出力するためのブールフラグ
func (b *Board) Print(force bool) {
	if !b.debug && !force {
		return
	}
	for i, c := range b.grid {
		fmt.Printf("%d ", c.answer)
		if i%9 == 8 {
			fmt.Print("\n")
		}
	}
}

// logはデバッグモードが有効な場合にフォーマットされたメッセージをログに記録します。
// フォーマット文字列とメッセージをフォーマットするための可変個の引数を受け取ります。
//
// パラメータ:
//   - format: ログメッセージのフォーマットを指定する文字列
//   - a: フォーマット文字列に従ってフォーマットされる可変個の引数
func (b *Board) log(format string, a ...any) {
	if !b.debug {
		return
	}
	fmt.Printf(format+"\n", a...)
}

// SolveLevel1は数独を解きます。
// これには次の2つの主要なステップを反復して実行します。
// 1. Cellに候補が1つしかない場合、その値で確定します。
// 2. 各グループ（行、列、3x3のマス目）について、候補の数字が1つのCellにしか入らない場合、そのCellをその値で確定します。
//
// パラメータ:
//   - depth: 再帰呼び出しの現在の深さを表す整数（この関数そのものは再起呼び出しになっていないがSolve2から呼び出された際のログの整形に使用）
//
// 戻り値:
//   - 数独が正常に解けたかどうかを示すブール値
//   - 数独が解けた場合は解決されたグリッドを表すCellのslice、解けなかった場合はnil
func (b *Board) SolveLevel1(depth int) (bool, []Cell) {
	groups := make([][]*Cell, 0, 27)
	groups = append(groups, b.rows...)
	groups = append(groups, b.collumns...)
	groups = append(groups, b.squares...)
	spaces := strings.Repeat("  ", depth)

	loop := 0
	for {
		loop++
		step1Loop := 0
		for {
			step1Loop++
			b.log("%sSTEP1 %d - %d", spaces, loop, step1Loop)
			step1Updated := false
			for i, c := range b.grid {
				if c.isFixed {
					continue
				}
				if len(c.candidates) == 1 {
					answer := c.candidates[0]
					b.Update(i, answer)
					b.log("%sPOS[%d] %d", spaces, i, answer)
					// b.Print(false)
					// b.log("")
					step1Updated = true
				}
			}
			if !step1Updated {
				break
			}
		}

		isUpdated := false
		b.log("%sSTEP2 %d", spaces, loop)
		for gi, g := range groups {
			for i := 1; i <= 9; i = i + 1 {
				var candidateCell *Cell
				candidateCount := 0
				for _, c := range g {
					if c.isFixed {
						continue
					}

					isCandidate := slices.Contains(c.candidates, i)
					if isCandidate {
						candidateCount++
						candidateCell = c
					}
					if candidateCount > 1 {
						break
					}
				}
				if candidateCount == 1 && candidateCell != nil {
					// for j, c := range g {
					// 	if !c.isFixed {
					// 		b.log("%sGROUP[%d] %d %d", spaces, gi, j, c.candidates)
					// 	}
					// }
					b.log("%sGROUP[%d] FIXED %d", spaces, gi, i)
					b.Update(candidateCell.pos, i)
					// b.Print(false)
					// b.log("")
					isUpdated = true
				}
			}
		}

		if !isUpdated {
			break
		}
	}

	isClear := true
	for _, c := range b.grid {
		if !c.isFixed {
			isClear = false
			break
		}
	}
	if isClear {
		return true, b.grid
	} else {
		return false, nil
	}
}

// SolveLevel2は数独を解きます。
// SolveLevel1では解けない場合、確定していないCellに対して候補から答えを仮定して再起的に解いていきます。
//
// パラメータ:
//   - depth: 再帰呼び出しの現在の深さを表す整数
//
// 戻り値:
//   - 数独が正常に解けたかどうかを示すブール値
//   - 数独が解けた場合は解決されたグリッドを表すCellのslice、解けなかった場合はnil
func (b *Board) SolveLevel2(depth int) (bool, []Cell) {
	spaces := strings.Repeat("  ", depth)
	for {
		isClear, _ := b.SolveLevel1(depth)
		if isClear {
			return true, b.grid
		}

		if b.noCandidates() {
			return false, nil
		}

		b.log("%sSTEP3", spaces)
		isUpdated := false
		for i, c := range b.grid {
			if c.isFixed {
				continue
			}
			tempCnadidates := make([]int, len(c.candidates))
			copy(tempCnadidates, c.candidates)
			for _, candidate := range tempCnadidates {
				b.log("%s深さ: %d POS[%d]を%dと仮定。候補%d", spaces, depth, i, candidate, tempCnadidates)
				temp := NewBoardWithGrid(b.grid, b.debug)
				temp.Update(i, candidate)
				solved, grid := temp.SolveLevel2(depth + 1)
				if solved {
					copy(b.grid, grid)
					b.log("%sクリア！深さ: %d POS[%d]を%dと仮定", spaces, depth, i, candidate)
					return true, grid
				} else {
					b.log("%s矛盾発生！深さ: %d POS[%d]を%dと仮定", spaces, depth, i, candidate)
					// b.log("%sremoveCandidate before %d %v %d", spaces, b.grid[i].pos, b.grid[i].candidates, candidate)
					b.grid[i].removeCandidate(candidate)
					// b.log("%sremoveCandidate after %d %v %d", spaces, b.grid[i].pos, b.grid[i].candidates, candidate)
					isUpdated = true
				}
			}
		}

		if !isUpdated {
			break
		}
	}

	isClear := true
	for _, c := range b.grid {
		if !c.isFixed {
			isClear = false
			break
		}
	}
	if isClear {
		return true, b.grid
	} else {
		return false, nil
	}
}

// noCandidatesは、盤面のグリッド内に固定されておらず、候補が残っていないセルがあるかどうかを確認します。
// そのようなセルが見つかった場合、これ以上の手が打てない矛盾した状態であることを示すためにtrueを返します。
func (b *Board) noCandidates() bool {
	for _, c := range b.grid {
		if !c.isFixed && len(c.candidates) == 0 {
			return true
		}
	}
	return false
}
