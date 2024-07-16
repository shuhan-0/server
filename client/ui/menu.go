package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

type Menu struct {
	options         []string
	selected        int
	backgroundWhite bool
	backgroundColor color.Color
}

func NewMenu(options []string) *Menu {
	return &Menu{
		options:         options,
		selected:        0,
		backgroundWhite: false,
		backgroundColor: color.RGBA{0, 0, 40, 255},
	}
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.selected = (m.selected + 1) % len(m.options)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.selected = (m.selected - 1 + len(m.options)) % len(m.options)
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(m.backgroundColor)

	if m.backgroundWhite {
		ebitenutil.DebugPrintAt(screen, "Background: Sky Blue", 100, 50)
	} else {
		ebitenutil.DebugPrintAt(screen, "Background: Dark Blue", 100, 50)
	}

	for i, option := range m.options {
		y := 100 + i*30
		if i == m.selected {
			ebitenutil.DebugPrintAt(screen, "> "+option, 100, y)
		} else {
			ebitenutil.DebugPrintAt(screen, "  "+option, 100, y)
		}
	}
}

func (m *Menu) GetSelected() string {
	return m.options[m.selected]
}

func (m *Menu) ToggleBackground() {
	m.backgroundWhite = !m.backgroundWhite
	if m.backgroundWhite {
		m.backgroundColor = color.RGBA{133, 176, 190, 255}
	} else {
		m.backgroundColor = color.RGBA{0, 0, 40, 255}
	}
}

func (m *Menu) GetBackgroundColor() color.Color {
	return m.backgroundColor
}
