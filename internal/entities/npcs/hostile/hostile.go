package hostile

import (
	"math"
	"math/rand"
	"time"

	"image/color"

	"github.com/coopstools/minibeast/internal/entities"
	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/models/properties"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type NPC struct {
	entities.BaseEntity
	Type           NPCType
	Stats          *models.Character
	DetectionRange float64
	State          MovementState
	BehaviorState  BehaviorState
	DamageAmount   int
}

type BehaviorState struct {
	WanderAngle float64
	StateTimer  float64
	LastUpdate  time.Time
	StartPos    properties.Position
}

func New(npcType NPCType, x, y float64) *NPC {
	npc := &NPC{
		BaseEntity: entities.BaseEntity{
			Position: properties.Position{X: x, Y: y},
			Size:     16.0,
		},
		Type:           npcType,
		DetectionRange: DefaultDetectionRange,
		State:          Idle,
		BehaviorState: BehaviorState{
			WanderAngle: rand.Float64() * math.Pi * 2,
			StateTimer:  rand.Float64()*(MaxWanderTime-MinWanderTime) + MinWanderTime,
			LastUpdate:  time.Now(),
			StartPos:    properties.Position{X: x, Y: y},
		},
		Stats:        models.NewCharacter(),
		DamageAmount: DefaultDamageAmount,
	}
	npc.Health = properties.NewHealth(npc.Stats.Vitality)
	return npc
}

func (n *NPC) Update(playerX, playerY float64, otherNPCs []*NPC) error {
	now := time.Now()
	deltaTime := now.Sub(n.BehaviorState.LastUpdate).Seconds()
	n.BehaviorState.LastUpdate = now

	// Calculate distance to player
	distToPlayer := math.Sqrt(math.Pow(n.Position.X-playerX, 2) + math.Pow(n.Position.Y-playerY, 2))

	// Update state based on player distance
	if distToPlayer <= n.DetectionRange {
		n.State = Pursuing
	} else if distToPlayer > n.DetectionRange*1.2 {
		if n.State == Pursuing {
			n.State = Idle
		}
	}

	// Calculate avoidance vector from other NPCs
	avoidX, avoidY := n.calculateAvoidance(otherNPCs)

	// Update behavior based on state
	switch n.State {
	case Pursuing:
		n.handlePursuit(playerX, playerY, avoidX, avoidY)
	case Wandering:
		n.handleWandering(deltaTime, avoidX, avoidY)
	case Pausing:
		n.handlePausing(deltaTime)
	case Idle:
		n.handleIdle(deltaTime)
	}

	n.UpdateMovement()
	return nil
}

func (n *NPC) UpdateMovement() {
	// Store previous position in case we need to revert
	prevX := n.Position.X
	prevY := n.Position.Y

	// Apply velocity to position
	n.Position.X += n.VelX
	n.Position.Y += n.VelY

	// Bound checking
	if n.Position.X < 0 || n.Position.X > 800-n.Size {
		// Revert X movement and bounce
		n.Position.X = prevX
		n.VelX = -n.VelX * 0.5 // Bounce with reduced velocity
	}
	if n.Position.Y < 0 || n.Position.Y > 600-n.Size {
		// Revert Y movement and bounce
		n.Position.Y = prevY
		n.VelY = -n.VelY * 0.5 // Bounce with reduced velocity
	}
}

func (n *NPC) calculateAvoidance(otherNPCs []*NPC) (float64, float64) {
	var avoidX, avoidY float64
	for _, other := range otherNPCs {
		if other == n {
			continue
		}

		dx := n.Position.X - other.Position.X
		dy := n.Position.Y - other.Position.Y
		dist := math.Sqrt(dx*dx + dy*dy)

		if dist < NPCAvoidRadius && dist > 0 {
			strength := (NPCAvoidRadius - dist) / NPCAvoidRadius * NPCAvoidStrength
			avoidX += (dx / dist) * strength
			avoidY += (dy / dist) * strength
		}
	}
	return avoidX, avoidY
}

func (n *NPC) handlePursuit(playerX, playerY, avoidX, avoidY float64) {
	dx := playerX - n.Position.X
	dy := playerY - n.Position.Y

	length := math.Sqrt(dx*dx + dy*dy)
	if length != 0 {
		dx /= length
		dy /= length
	}

	// Add avoidance vector
	dx += avoidX
	dy += avoidY

	length = math.Sqrt(dx*dx + dy*dy)
	if length != 0 {
		dx /= length
		dy /= length
	}

	n.VelX = dx * MaxSpeed * 0.7
	n.VelY = dy * MaxSpeed * 0.7
}

func (n *NPC) handleWandering(deltaTime float64, avoidX, avoidY float64) {
	n.BehaviorState.StateTimer -= deltaTime
	if n.BehaviorState.StateTimer <= 0 {
		n.State = Pausing
		n.BehaviorState.StateTimer = rand.Float64()*(MaxPauseTime-MinPauseTime) + MinPauseTime
		n.VelX = 0
		n.VelY = 0
		return
	}

	// Check if too far from start position
	distFromStart := math.Sqrt(
		math.Pow(n.Position.X-n.BehaviorState.StartPos.X, 2) +
			math.Pow(n.Position.Y-n.BehaviorState.StartPos.Y, 2))

	if distFromStart > WanderRadius {
		dx := n.BehaviorState.StartPos.X - n.Position.X
		dy := n.BehaviorState.StartPos.Y - n.Position.Y
		length := math.Sqrt(dx*dx + dy*dy)
		if length != 0 {
			n.BehaviorState.WanderAngle = math.Atan2(dy, dx)
		}
	} else if rand.Float64() < 0.02 { // 2% chance per update
		n.BehaviorState.WanderAngle += (rand.Float64() - 0.5) * math.Pi / 2
	}

	dx := math.Cos(n.BehaviorState.WanderAngle)
	dy := math.Sin(n.BehaviorState.WanderAngle)

	// Add avoidance vector
	dx += avoidX
	dy += avoidY

	length := math.Sqrt(dx*dx + dy*dy)
	if length != 0 {
		dx /= length
		dy /= length
	}

	n.VelX = dx * WanderSpeed
	n.VelY = dy * WanderSpeed
}

func (n *NPC) handlePausing(deltaTime float64) {
	n.BehaviorState.StateTimer -= deltaTime
	if n.BehaviorState.StateTimer <= 0 {
		n.State = Wandering
		n.BehaviorState.StateTimer = rand.Float64()*(MaxWanderTime-MinWanderTime) + MinWanderTime
		n.BehaviorState.WanderAngle += (rand.Float64() - 0.5) * math.Pi
	}
	n.VelX = 0
	n.VelY = 0
}

func (n *NPC) handleIdle(deltaTime float64) {
	n.BehaviorState.StateTimer -= deltaTime
	if n.BehaviorState.StateTimer <= 0 {
		n.State = Wandering
		n.BehaviorState.StateTimer = rand.Float64()*(MaxWanderTime-MinWanderTime) + MinWanderTime
		n.BehaviorState.WanderAngle = rand.Float64() * math.Pi * 2
	}
	n.VelX = 0
	n.VelY = 0
}

func (n *NPC) Draw(screen *ebiten.Image) {
	var npcColor color.RGBA
	switch n.State {
	case Pursuing:
		npcColor = color.RGBA{160, 0, 160, 255}
	case Wandering:
		npcColor = color.RGBA{128, 0, 128, 255}
	case Pausing:
		npcColor = color.RGBA{100, 0, 100, 255}
	case Idle:
		npcColor = color.RGBA{80, 0, 80, 255}
	}

	vector.DrawFilledCircle(screen,
		float32(n.Position.X+n.Size/2),
		float32(n.Position.Y+n.Size/2),
		float32(n.Size/2),
		npcColor,
		false)
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

// Move NPC-specific methods here
