package ui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type HealthDisplay struct {
	X, Y     int
	TextFace *text.GoXFace
}

func NewHealthDisplay(x, y int, face font.Face) *HealthDisplay {
	return &HealthDisplay{
		X:        x,
		Y:        y,
		TextFace: text.NewGoXFace(face),
	}
}

func (h *HealthDisplay) Draw(screen *ebiten.Image, current, max int) {
	healthText := fmt.Sprintf("%d/%d", current, max)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(h.X), float64(h.Y))
	text.Draw(screen, healthText, h.TextFace, opts)
}
