package ui

import (
	"image/color"
	"strings"

	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Popup struct {
	X, Y          int
	Width, Height int
	Content       string
	Visible       bool
}

func NewPopup(width, height int) *Popup {
	return &Popup{
		Width:   width,
		Height:  height,
		Visible: false,
	}
}

func (p *Popup) Show(x, y int, content string) {
	p.X = x
	p.Y = y
	p.Content = content
	p.Visible = true
}

func (p *Popup) Hide() {
	p.Visible = false
}

func (p *Popup) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	if !p.Visible {
		return
	}

	// Draw semi-transparent background
	vector.DrawFilledRect(screen,
		float32(p.X), float32(p.Y),
		float32(p.Width), float32(p.Height),
		color.RGBA{0, 0, 0, 180}, // Dark semi-transparent background
		false)

	// Draw border
	vector.StrokeRect(screen,
		float32(p.X), float32(p.Y),
		float32(p.Width), float32(p.Height),
		1, color.White, false)

	// Draw content
	content := strings.Split(p.Content, "\n")
	xfont := text.NewGoXFace(assets.Font)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(p.X+10), float64(p.Y-10))
	for _, line := range content {
		opts.GeoM.Translate(0, 20)
		text.Draw(screen, line, xfont, opts)
	}
}
