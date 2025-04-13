package models

import (
	"math"

	"github.com/coopstools/minibeast/internal/models/properties"
)

type Entity interface {
	GetPosition() (float64, float64)
	Update(playerX, playerY float64)
	IsColliding(otherX, otherY float64) bool
}

// Common position and movement logic for all entities
type BaseEntity struct {
	Position   properties.Position
	VelX, VelY float64
	TargetVelX float64
	TargetVelY float64
	Size       float64
}

func (e *BaseEntity) GetPosition() (float64, float64) {
	return e.Position.X, e.Position.Y
}

func (e *BaseEntity) IsColliding(otherX, otherY float64) bool {
	dist := math.Sqrt(math.Pow(e.Position.X-otherX, 2) + math.Pow(e.Position.Y-otherY, 2))
	return dist < e.Size
}

func (e *BaseEntity) UpdateMovement() {
	// Update X velocity
	if e.VelX < e.TargetVelX {
		e.VelX = min(e.VelX+Acceleration, e.TargetVelX)
	} else if e.VelX > e.TargetVelX {
		e.VelX = max(e.VelX-Deceleration, e.TargetVelX)
	}

	// Update Y velocity
	if e.VelY < e.TargetVelY {
		e.VelY = min(e.VelY+Acceleration, e.TargetVelY)
	} else if e.VelY > e.TargetVelY {
		e.VelY = max(e.VelY-Deceleration, e.TargetVelY)
	}

	// Apply velocity
	e.Position.X += e.VelX
	e.Position.Y += e.VelY

	// Bound checking
	e.Position.X = max(0, min(e.Position.X, 800-e.Size))
	e.Position.Y = max(0, min(e.Position.Y, 600-e.Size))
}
