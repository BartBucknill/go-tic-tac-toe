package main

import (
	"os"

	game "github.com/BartBucknill/go-tic-tac-toe/game"
)

func main() {
	game.New(os.Stdout, os.Stdin, os.Stderr).Play()
}
