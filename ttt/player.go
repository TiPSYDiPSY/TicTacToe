package ttt

type Player interface {
	GetMove(*board) (int, int, error)
	Mark() string
	Name() string
}