package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

type Button struct {
	X, Y          int
	Width, Height int
	Text          string
	Hovered       bool
	TextFace      *text.GoXFace
}

func NewButton(x, y, width, height int, content string, face font.Face) *Button {
	return &Button{
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Text:     content,
		TextFace: text.NewGoXFace(face),
	}
}

func (b *Button) Contains(x, y int) bool {
	return x >= b.X && x < b.X+b.Width &&
		y >= b.Y && y < b.Y+b.Height
}

func (b *Button) Draw(screen *ebiten.Image) {
	bgColor := color.RGBA{60, 60, 60, 255}
	if b.Hovered {
		bgColor = color.RGBA{80, 80, 80, 255}
	}

	// Draw button background
	vector.DrawFilledRect(screen,
		float32(b.X), float32(b.Y),
		float32(b.Width), float32(b.Height),
		bgColor, false)

	// Draw button border
	vector.StrokeRect(screen,
		float32(b.X), float32(b.Y),
		float32(b.Width), float32(b.Height),
		1, color.White, false)

	// Center text in button
	textX := b.X + 2
	textY := b.Y + 5

	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(textX), float64(textY))
	text.Draw(screen, b.Text, b.TextFace, op)
}
