package models

type Node struct {
	Row    int
	Column int
	Next   *Node
	Prev   *Node
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func NewSnake() *List {
	node := &Node{
		Row:    0,
		Column: 0,
	}
	return &List{
		Head: node,
		Tail: node,
		Size: 1,
	}
}
