package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/sh-seike/sudoku-solver-go/sudoku"
)

func main() {
	debug := flag.Bool("d", false, "debug mode")
	level := flag.Int("l", 1, "level")
	flag.Parse()
	// fmt.Println(flag.Args())
	// fmt.Println(*debug)
	// fmt.Println(*level)
	// argsが空でない場合はfileNameにargs[1]を代入する。空の場合はquestion1.csvを代入する
	fileName := "question1.csv"
	if len(flag.Args()) > 0 {
		fileName = flag.Args()[0]
	}
	fmt.Println(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		file.Close()
		os.Exit(1)
	}

	board := sudoku.NewBoard(*debug)
	for y, r := range rows {
		for x, c := range r {
			i, err := strconv.Atoi(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				file.Close()
				os.Exit(1)
			}
			if i != 0 {
				board.Update(y*9+x, i)
			}
		}
	}
	board.Print(true)
	fmt.Println("")

	isClear := false
	if *level > 1 {
		isClear, _ = board.Solve2(0)
	} else {
		isClear, _ = board.Solve(0)
	}
	if isClear {
		fmt.Println("DONE!")
	} else {
		fmt.Println("GIVE UP!")
	}

	board.Print(true)
	fmt.Println("")
}
