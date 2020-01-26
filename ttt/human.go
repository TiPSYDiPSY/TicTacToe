package ttt

import (
	"errors"
	"fmt"
)

func NewHumanPlayer(n string, m string) *HumanPlayer {
	return  &HumanPlayer{
		name: n,
		mark: m,
	}
}

type HumanPlayer struct {
	name string
	mark string
}

// GetMove returns next move
func (p *HumanPlayer) GetMove(b *board) (int, int, error) {
	fmt.Println("Enter position: ")
	var i,j int

	if n, err := fmt.Scanf("%d %d", &i, &j); err != nil || n != 2 {
		return 0, 0, err
	}

	if i > 2 || j > 2 {
		err := errors.New("position can be only 0,1,2")
		return  0, 0, err
	}

	fmt.Println("You input:", i, j)
	return i, j, nil
}

// Mark is a getter for player's mark
func (p *HumanPlayer) Mark() string {
	return p.mark
}

// Name is a getter for player's name
func (p *HumanPlayer) Name() string {
	return p.name
}