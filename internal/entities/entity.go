package entities

import (
	"github.com/coopstools/minibeast/internal/models/properties"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Update() error
	Draw(screen *ebiten.Image)
	GetPosition() (float64, float64)
	IsColliding(otherX, otherY float64) bool
}

// Common base functionality
type BaseEntity struct {
	Position   properties.Position
	Size       float64
	VelX, VelY float64
	Health     *properties.Health
}
