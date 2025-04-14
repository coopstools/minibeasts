package models

import (
	"math"
	"time"

	"github.com/coopstools/minibeast/internal/models/properties"
)

type PlayerPosition struct {
	Position   properties.Position
	VelX, VelY float64 // Current velocity
	TargetVelX float64 // Target velocity based on input
	TargetVelY float64
	Knockback  struct { // New knockback system
		Active     bool
		VelX, VelY float64
		Duration   time.Duration
		StartTime  time.Time
	}
}

const (
	PlayerSize        = 16.0
	MaxSpeed          = 4.0                    // Maximum movement speed
	Acceleration      = 0.5                    // How quickly we reach max speed
	Deceleration      = 0.3                    // How quickly we slow down
	KnockbackForce    = 8.0                    // Initial knockback velocity
	KnockbackDuration = time.Millisecond * 200 // How long knockback lasts
)

// Helper function to apply smooth movement
func (p *PlayerPosition) UpdateMovement() {
	now := time.Now()

	if p.Knockback.Active {
		// Check if knockback has expired
		if now.Sub(p.Knockback.StartTime) >= p.Knockback.Duration {
			p.Knockback.Active = false
		} else {
			// Apply knockback movement
			p.Position.X += p.Knockback.VelX
			p.Position.Y += p.Knockback.VelY

			// Gradually reduce knockback velocity
			reduction := 1.0 - (now.Sub(p.Knockback.StartTime).Seconds() / p.Knockback.Duration.Seconds())
			p.Knockback.VelX *= reduction
			p.Knockback.VelY *= reduction

			// Skip normal movement while in knockback
			// Bound checking
			p.Position.X = max(0, min(p.Position.X, 800-PlayerSize))
			p.Position.Y = max(0, min(p.Position.Y, 600-PlayerSize))
			return
		}
	}

	// Normal movement code
	if p.VelX < p.TargetVelX {
		p.VelX = min(p.VelX+Acceleration, p.TargetVelX)
	} else if p.VelX > p.TargetVelX {
		p.VelX = max(p.VelX-Deceleration, p.TargetVelX)
	}

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

func (p *PlayerPosition) ApplyKnockback(sourceX, sourceY float64) {
	// Calculate direction away from the source
	dx := p.Position.X - sourceX
	dy := p.Position.Y - sourceY

	// Normalize the direction
	length := math.Sqrt(dx*dx + dy*dy)
	if length > 0 {
		dx /= length
		dy /= length
	}

	// Apply knockback
	p.Knockback.Active = true
	p.Knockback.VelX = dx * KnockbackForce
	p.Knockback.VelY = dy * KnockbackForce
	p.Knockback.StartTime = time.Now()
	p.Knockback.Duration = KnockbackDuration
}
