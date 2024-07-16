package network

import (
	"encoding/json"
	"net"
	"snake-game/internal/game"
	"snake-game/pkg/protocol"
	"sync"
)

type Server struct {
	game     *game.Game
	listener net.Listener
	clients  map[string]net.Conn
	mutex    sync.Mutex
}

func NewServer(game *game.Game, address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{
		game:     game,
		listener: listener,
		clients:  make(map[string]net.Conn),
	}, nil
}

func (s *Server) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var loginMsg protocol.Message
	err := decoder.Decode(&loginMsg)
	if err != nil || loginMsg.Type != protocol.Login {
		return
	}

	loginPayload, ok := loginMsg.Payload.(map[string]interface{})
	if !ok {
		return
	}
	playerID, ok := loginPayload["username"].(string)
	if !ok {
		return
	}

	s.mutex.Lock()
	s.clients[playerID] = conn
	s.game.AddPlayer(playerID, "PlayerName") // Example: replace with actual player name
	s.mutex.Unlock()

	defer func() {
		s.mutex.Lock()
		delete(s.clients, playerID)
		s.game.RemovePlayer(playerID)
		s.mutex.Unlock()
	}()

	for {
		var msg protocol.Message
		err := decoder.Decode(&msg)
		if err != nil {
			break
		}

		s.mutex.Lock()
		switch msg.Type {
		case protocol.Move:
			movePayload, ok := msg.Payload.(map[string]interface{})
			if !ok {
				continue
			}
			directionStr, ok := movePayload["direction"].(string)
			if !ok {
				continue
			}
			var direction game.Direction
			switch directionStr {
			case "Up":
				direction = game.Up
			case "Down":
				direction = game.Down
			case "Left":
				direction = game.Left
			case "Right":
				direction = game.Right
			default:
				continue
			}
			newPosition := calculateNewPosition(direction) // Replace with your actual logic to calculate new position
			s.game.MovePlayer(playerID, newPosition)
		}
		gameState := s.game.GetState()
		s.mutex.Unlock()

		s.mutex.Lock()
		err = encoder.Encode(protocol.Message{
			Type:    protocol.GameState,
			Payload: gameState,
		})
		s.mutex.Unlock()

		if err != nil {
			break
		}
	}
}

func (s *Server) Broadcast(msg protocol.Message) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, conn := range s.clients {
		json.NewEncoder(conn).Encode(msg)
	}
}

// Helper function to calculate new position based on direction
func calculateNewPosition(direction game.Direction) game.Position {
	// Replace with actual logic to calculate new position based on direction
	// This is just a placeholder
	return game.Position{X: 1, Y: 1}
}
