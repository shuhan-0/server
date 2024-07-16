package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"snake-game/client/network"
	"snake-game/internal/game"
)

type GameWindow struct {
	client          *network.Client
	width           int
	height          int
	game            *game.Game
	menu            *Menu
	paused          bool
	backgroundColor color.Color
}

func NewGameWindow(client *network.Client, menu *Menu) *GameWindow {
	return &GameWindow{
		client:          client,
		width:           800,
		height:          600,
		game:            game.NewGame(40, 30),
		menu:            menu,
		backgroundColor: color.RGBA{0, 0, 40, 255},
	}
}

func (gw *GameWindow) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		gw.paused = !gw.paused
	}

	if !gw.paused && !gw.game.IsGameOver() {
		gw.game.Update()

		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			gw.game.ChangeDirection(game.Up)
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			gw.game.ChangeDirection(game.Down)
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			gw.game.ChangeDirection(game.Left)
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			gw.game.ChangeDirection(game.Right)
		}
	}

	return nil
}

func (gw *GameWindow) Draw(screen *ebiten.Image) {
	screen.Fill(gw.backgroundColor)

	borderWidth := 3
	mapWidth := gw.game.Width * 20
	mapHeight := gw.game.Height * 20

	ebitenutil.DrawRect(screen, 0, 0, float64(mapWidth), float64(borderWidth), color.White)
	ebitenutil.DrawRect(screen, 0, float64(mapHeight)-float64(borderWidth), float64(mapWidth), float64(borderWidth), color.White)
	ebitenutil.DrawRect(screen, 0, 0, float64(borderWidth), float64(mapHeight), color.White)
	ebitenutil.DrawRect(screen, float64(mapWidth)-float64(borderWidth), 0, float64(borderWidth), float64(mapHeight), color.White)

	if gw.game.IsGameOver() {
		ebitenutil.DebugPrintAt(screen, "Game Over", 400, 250)
		ebitenutil.DebugPrintAt(screen, "Press Y to restart, N to exit to main menu", 300, 280)
	} else {
		for i, p := range gw.game.Snake.Body {
			if i == 0 {
				ebitenutil.DrawRect(screen, float64(p.X*20), float64(p.Y*20), 20, 20, color.RGBA{0, 128, 0, 255})
			} else {
				ebitenutil.DrawRect(screen, float64(p.X*20), float64(p.Y*20), 20, 20, color.RGBA{85, 107, 47, 255})
			}
		}

		foodPos := gw.game.Food
		ebitenutil.DrawCircle(screen, float64(foodPos.X*20)+10, float64(foodPos.Y*20)+10, 8, color.RGBA{255, 255, 0, 255})
	}

	if gw.paused {
		ebitenutil.DebugPrintAt(screen, "PAUSED", 400, 300)
		ebitenutil.DebugPrintAt(screen, "Press P to resume", 380, 330)
	}
}
func (gw *GameWindow) getTextColor() color.Color {
	r, g, b, _ := gw.backgroundColor.RGBA()
	if (r+g+b)/3 > 32768 { // 如果背景颜色较浅
		return color.Black
	}
	return color.White
}

func (gw *GameWindow) Layout(outsideWidth, outsideHeight int) (int, int) {
	return gw.width, gw.height
}

func (gw *GameWindow) RestartGame() {
	gw.game = game.NewGame(40, 30)
	gw.paused = false
}

func (gw *GameWindow) IsGameOver() bool {
	return gw.game.IsGameOver()
}

func (gw *GameWindow) SetBackgroundColor(c color.Color) {
	gw.backgroundColor = c
}
