package main

import (
	"flag"
	"log"
	"snake-game/client/network"
	"snake-game/client/ui"
	"snake-game/pkg/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client, err := network.NewClient(cfg.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	menu := ui.NewMenu([]string{"Start Game", "Change Background", "Developer Info", "Quit Game"})
	gameWindow := ui.NewGameWindow(client, menu)
	game := ui.NewGame(gameWindow, menu, client)

	if err := game.Run(); err != nil {
		log.Fatalf("Game error: %v", err)
	}
}
