package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrorBounds = errors.New("out of bounds")
	ErrorDigit  = errors.New("invalid digit")
)

const rows, columns = 9, 9

type Grid [rows][columns]int8

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	} else if column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError

	if !inBounds(row, column) {
		errs = append(errs, ErrorBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrorDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func main() {
	var g Grid
	err := g.Set(10, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Println(errs, ok)
			fmt.Printf("%d error(s) occurred:\n", len(errs))

			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}
}
