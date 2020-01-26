package main

import "TicTacToe/ttt"

func main() {
	hp := ttt.NewHumanPlayer("Deniss", "X")
	cp := ttt.NewHumanPlayer("AI", "0")
	g := ttt.NewGame(hp, cp)
	g.Start()
}