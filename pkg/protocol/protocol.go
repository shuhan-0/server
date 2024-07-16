package protocol

type Message struct {
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}

type MessageType int

const (
	Login MessageType = iota
	Move
	GameState
	// Add more message types as needed
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerState struct {
	// Define player state structure
}

type GameStatePayload struct {
	Players []PlayerState `json:"players"`
	Food    []Position    `json:"food"`
}

func PositionFromString(str string) Position {
	// Implement parsing string to Position struct
	return Position{}
}
