package ui

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"snake-game/client/network"
	"time"
)

type Game struct {
	window                   *GameWindow
	menu                     *Menu
	client                   *network.Client
	state                    GameState
	lastBackgroundChangeTime time.Time
}

type GameState int

const (
	MenuState GameState = iota
	PlayingState
	DeveloperInfoState
)

var ErrGameTermination = errors.New("game termination")

func NewGame(window *GameWindow, menu *Menu, client *network.Client) *Game {
	return &Game{
		window:                   window,
		menu:                     menu,
		client:                   client,
		state:                    MenuState,
		lastBackgroundChangeTime: time.Now(),
	}
}

func (g *Game) Update() error {
	switch g.state {
	case MenuState:
		g.menu.Update()
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			switch g.menu.GetSelected() {
			case "Start Game":
				g.state = PlayingState
				g.window.RestartGame()
			case "Change Background":
				currentTime := time.Now()
				if currentTime.Sub(g.lastBackgroundChangeTime) >= 500*time.Millisecond {
					g.menu.ToggleBackground()
					g.window.SetBackgroundColor(g.menu.GetBackgroundColor())
					g.lastBackgroundChangeTime = currentTime
				}
			case "Developer Info":
				g.state = DeveloperInfoState
			case "Quit Game":
				return ErrGameTermination
			}
		}
	case PlayingState:
		if g.window.IsGameOver() {
			if ebiten.IsKeyPressed(ebiten.KeyY) {
				g.window.RestartGame()
			} else if ebiten.IsKeyPressed(ebiten.KeyN) {
				g.state = MenuState
			}
		} else {
			return g.window.Update()
		}
	case DeveloperInfoState:
		if ebiten.IsKeyPressed(ebiten.KeyEscape) {
			g.state = MenuState
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case MenuState:
		g.menu.Draw(screen)
	case PlayingState:
		g.window.Draw(screen)
	case DeveloperInfoState:
		g.drawDeveloperInfo(screen)
	}
}

func (g *Game) drawDeveloperInfo(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "shuhan-0", 100, 100)
	ebitenutil.DebugPrintAt(screen, "Version: 1.0", 100, 130)
	ebitenutil.DebugPrintAt(screen, "github：https://github.com/shuhan-0/", 100, 160)
	ebitenutil.DebugPrintAt(screen, "Press ESC to return to menu", 100, 190)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.window.Layout(outsideWidth, outsideHeight)
}

func (g *Game) Run() error {
	ebiten.SetWindowSize(g.window.width, g.window.height)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(g); err != nil && !errors.Is(err, ErrGameTermination) {
		return err
	}
	return nil
}
