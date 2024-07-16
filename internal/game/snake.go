package game

type Snake struct {
	Body      []Position
	Direction Direction
	grow      bool
}

func NewSnake(startPos Position) *Snake {
	return &Snake{
		Body:      []Position{startPos},
		Direction: Right,
	}
}

func (s *Snake) Head() Position {
	return s.Body[0]
}

func (s *Snake) Move() {
	head := s.Head()
	newHead := head

	switch s.Direction {
	case Up:
		newHead.Y--
	case Down:
		newHead.Y++
	case Left:
		newHead.X--
	case Right:
		newHead.X++
	}

	s.Body = append([]Position{newHead}, s.Body...)
	if !s.grow {
		s.Body = s.Body[:len(s.Body)-1]
	} else {
		s.grow = false
	}
}

func (s *Snake) Grow() {
	s.grow = true
}

func (s *Snake) ChangeDirection(dir Direction) {
	if (s.Direction == Up && dir != Down) ||
		(s.Direction == Down && dir != Up) ||
		(s.Direction == Left && dir != Right) ||
		(s.Direction == Right && dir != Left) {
		s.Direction = dir
	}
}
