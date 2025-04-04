package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// PlaceholderSprite represents a simple colored square
type PlaceholderSprite struct {
	Width  int
	Height int
	Color  color.Color
}

func NewPlaceholderSprite(width, height int, col color.Color) *PlaceholderSprite {
	return &PlaceholderSprite{
		Width:  width,
		Height: height,
		Color:  col,
	}
}

func (p *PlaceholderSprite) Draw(screen *ebiten.Image, x, y float32) {
	vector.DrawFilledRect(screen, 2, 2, float32(p.Width), float32(p.Height), p.Color, false)
}

// Common placeholder colors
var (
	PlayerColor   = color.RGBA{255, 0, 0, 255}     // Red
	NPCColor      = color.RGBA{0, 255, 0, 255}     // Green
	ObstacleColor = color.RGBA{128, 128, 128, 255} // Gray
	InteractColor = color.RGBA{255, 255, 0, 255}   // Yellow
)

// Helper functions for common game elements
func NewPlayerPlaceholder() *PlaceholderSprite {
	return NewPlaceholderSprite(14, 14, PlayerColor)
}

func NewNPCPlaceholder() *PlaceholderSprite {
	return NewPlaceholderSprite(14, 14, NPCColor)
}

func NewObstaclePlaceholder() *PlaceholderSprite {
	return NewPlaceholderSprite(14, 14, ObstacleColor)
}

func NewInteractivePlaceholder() *PlaceholderSprite {
	return NewPlaceholderSprite(14, 14, InteractColor)
}
