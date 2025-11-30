package models

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type Node struct {
	Row    int
	Column int
	Next   *Node
	Prev   *Node
}

type Snake struct {
	Head      *Node
	Tail      *Node
	Direction Direction
	Size      int
}

func NewSnake() *Snake {
	node := &Node{
		Row:    0,
		Column: 0,
	}
	return &Snake{
		Head:      node,
		Tail:      node,
		Direction: Right,
		Size:      1,
	}
}

func (s *Snake) AddToHead() error {
	newHead := &Node{
		Row:    s.Head.Row,
		Column: s.Head.Column,
		Next:   s.Head,
	}

	switch s.Direction {
	case Up:
		newHead.Row--
	case Down:
		newHead.Row++
	case Left:
		newHead.Column--
	case Right:
		newHead.Column++
	}

	newHead.Prev = s.Head
	s.Head.Next = newHead
	s.Head = newHead
	s.Size++
	return nil
}

func (s *Snake) RemoveFromTail() error {
	if s.Size <= 1 {
		return nil
	}
	newTail := s.Tail.Next
	newTail.Prev = nil
	s.Tail = newTail
	s.Size--
	return nil
}

func (s *Snake) GetHeadLocation() (int, int) {
	return s.Head.Row, s.Head.Column
}

func (s *Snake) SetDirection(dir Direction) {
	s.Direction = dir
}
