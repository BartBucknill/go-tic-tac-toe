package game

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Game struct {
	board board
	human rune
	out   io.Writer
	in    io.Reader
	err   io.Writer
}

type coordinates map[string]int

type board [][]rune

func New(out io.Writer, in io.Reader, err io.Writer) Game {
	return Game{
		out:   out,
		in:    in,
		err:   err,
		human: 'O',
		board: board{
			{' ', '1', '2', '3'},
			{'A', '.', '.', '.'},
			{'B', '.', '.', '.'},
			{'C', '.', '.', '.'},
		}}
}

var yCoords = coordinates{
	"a": 1,
	"b": 2,
	"c": 3,
}

var xCoords = coordinates{
	"1": 1,
	"2": 2,
	"3": 3,
}

func updateBoard(symbol rune, cmd string, b board) error {

	coordCodes := strings.Split(strings.ToLower(strings.TrimSpace(strings.TrimSuffix(cmd, "\n"))), "")
	var x, y int
	if len(coordCodes) > 2 {
		return fmt.Errorf("received more than two coordinates! %v", coordCodes)
	}
	for _, code := range coordCodes {
		if coord, ok := xCoords[code]; ok {
			x = coord
		} else if coord, ok := yCoords[code]; ok {
			y = coord
		}
	}
	if x == 0 || y == 0 {
		return fmt.Errorf("unable to interpret coordinates: %v", coordCodes)
	}

	value := b[y][x]
	if value != '.' {
		return fmt.Errorf("position %v unavailable", cmd)
	}
	b[y][x] = symbol
	return nil

}

func (g Game) render() {
	fmt.Fprintf(g.out, "\n")
	for _, row := range g.board {
		for _, r := range row {
			fmt.Fprintf(g.out, "%s  ", string(r))
		}
		fmt.Fprintf(g.out, "\n\n")
	}
}

func (g Game) Play() {
	for {

		g.render()
		fmt.Fprintf(g.out, "> ")
		reader := bufio.NewReader(g.in)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(g.err, "An error occured while reading input. Please try again.\n%v", err)
			return
		}
		if err := updateBoard(g.human, input, g.board); err != nil {
			fmt.Fprintf(g.err, "%v\n", err)
			return
		}

		g.render()
	}

}
