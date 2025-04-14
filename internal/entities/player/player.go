package player

import (
	"time"

	"github.com/coopstools/minibeast/internal/entities"
	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/models/properties"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	entities.BaseEntity
	Character  *models.Character
	Knockback  KnockbackState
	InputState InputState
}

type KnockbackState struct {
	Active     bool
	VelX, VelY float64
	Duration   time.Duration
	StartTime  time.Time
}

type InputState struct {
	TargetVelX float64
	TargetVelY float64
}

func New(character *models.Character) *Player {
	return &Player{
		BaseEntity: entities.BaseEntity{
			Position: properties.Position{X: 400, Y: 300},
			Size:     16.0,
			Health:   character.Health,
		},
		Character: character,
		Knockback: KnockbackState{
			Active:   false,
			VelX:     0,
			VelY:     0,
			Duration: time.Second / 2,
		},
	}
}

func (p *Player) Update() error {
	// Handle keyboard input for movement
	p.InputState.TargetVelX = 0
	p.InputState.TargetVelY = 0

	speed := 2.0
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		speed = 3.0
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.InputState.TargetVelY = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.InputState.TargetVelY = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.InputState.TargetVelX = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.InputState.TargetVelX = speed
	}

	// Update position based on velocity
	p.Position.X += p.InputState.TargetVelX
	p.Position.Y += p.InputState.TargetVelY

	// Bound checking
	p.Position.X = max(0, min(p.Position.X, 800-p.Size))
	p.Position.Y = max(0, min(p.Position.Y, 600-p.Size))

	return nil
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
