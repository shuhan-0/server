package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"snake-game/internal/auth"
	"snake-game/internal/cache"
	"snake-game/internal/database"
	"snake-game/pkg/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewMySQLDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient, err := cache.NewRedisClient(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	authService := auth.NewAuthService(db, redisClient)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := authService.Authenticate(loginReq.Username, loginReq.Password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	})

	log.Printf("Starting login gate on %s", cfg.LoginGateAddress)
	if err := http.ListenAndServe(cfg.LoginGateAddress, nil); err != nil {
		log.Fatalf("Login gate error: %v", err)
	}
}
