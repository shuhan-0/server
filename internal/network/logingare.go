package network

import (
	"encoding/json"
	"log"
	"net"
	"snake-game/internal/auth"
)

type LoginGate struct {
	address     string
	authService *auth.AuthService
}

func NewLoginGate(address string, authService *auth.AuthService) *LoginGate {
	return &LoginGate{
		address:     address,
		authService: authService,
	}
}

func (lg *LoginGate) Start() error {
	listener, err := net.Listen("tcp", lg.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go lg.handleConnection(conn)
	}
}

func (lg *LoginGate) handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := decoder.Decode(&loginRequest); err != nil {
		log.Printf("Error decoding login request: %v", err)
		return
	}

	token, err := lg.authService.Authenticate(loginRequest.Username, loginRequest.Password)
	if err != nil {
		encoder.Encode(map[string]string{"error": "Authentication failed"})
		return
	}

	encoder.Encode(map[string]string{"token": token})
}
