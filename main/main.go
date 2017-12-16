package main

import (
	"fmt"

	g "tictactoe/game"
)

func main() {
	game, err := g.NewGame(3, 2, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	game.Loop()
}
