package ttt

import  . "fmt"

type Game struct {
	p1 Player
	p2 Player
	current Player
	board *board
	round int
}

func NewGame(p1 Player, p2 Player) *Game{
	return &Game{
		p1:      p1,
		p2:      p2,
		current: p1,
		board:   newBoard(),
		round:   1,
	}
}

func (g *Game) isOver() bool {
	return g.board.winner() != "" || g.board.emptyCount() == 0
}

func (g *Game) Start() {
	Println(" ")
	for !g.isOver() {
		g.printInfo()
		i, j , err := g.current.GetMove(g.board)

		if err != nil {
			Println("your input is invalid, please try again")
			continue
		}

		if g.board[i][j] != "_" {
			Println("position is already occupied, please try again")
			continue
		}

		g.board[i][j] = g.current.Mark()
		g.switchPlayer()
		g.round++
	}

	Println(g.board)
	Println("Game Over!")
}

func (g *Game) switchPlayer() {
	if g.current == g.p1 {
		g.current = g.p2
	} else {
		g.current = g.p1
	}
}

func (g *Game) printInfo() {
	Printf("Turn %d\n", g.round)
	Println(g.board)
	Println("Current player:", g.current.Name())
}