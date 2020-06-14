package main

import (
	"errors"
	"fmt"
	"os"
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

func (g *Grid) Set(row, column int, digit int8) error {
	if !inBounds(row, column) {
		return ErrorBounds
	}
	if !validDigit(digit) {
		return ErrorDigit
	}
	g[row][column] = digit
	return nil
}

func main() {
	var g Grid
	err := g.Set(10, 0, 5)
	if err != nil {
		switch err {
		case ErrorBounds, ErrorDigit:
			fmt.Println("Les erreurs de paramètres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}
}
