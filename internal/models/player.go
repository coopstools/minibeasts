package models

import "github.com/coopstools/minibeast/internal/models/properties"

type PlayerPosition struct {
	Position   properties.Position
	VelX, VelY float64 // Current velocity
	TargetVelX float64 // Target velocity based on input
	TargetVelY float64
}

const (
	PlayerSize   = 16.0
	MaxSpeed     = 4.0 // Maximum movement speed
	Acceleration = 0.5 // How quickly we reach max speed
	Deceleration = 0.3 // How quickly we slow down
)

// Helper function to apply smooth movement
func (p *PlayerPosition) UpdateMovement() {
	// Update X velocity
	if p.VelX < p.TargetVelX {
		p.VelX = min(p.VelX+Acceleration, p.TargetVelX)
	} else if p.VelX > p.TargetVelX {
		p.VelX = max(p.VelX-Deceleration, p.TargetVelX)
	}

	// Update Y velocity
	if p.VelY < p.TargetVelY {
		p.VelY = min(p.VelY+Acceleration, p.TargetVelY)
	} else if p.VelY > p.TargetVelY {
		p.VelY = max(p.VelY-Deceleration, p.TargetVelY)
	}

	// Apply velocity
	p.Position.X += p.VelX
	p.Position.Y += p.VelY

	// Bound checking
	p.Position.X = max(0, min(p.Position.X, 800-PlayerSize))
	p.Position.Y = max(0, min(p.Position.Y, 600-PlayerSize))
}
