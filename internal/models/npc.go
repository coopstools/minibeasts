package models

import (
	"math"
	"math/rand"
	"time"

	"github.com/coopstools/minibeast/internal/models/properties"
)

type NPCType int

const (
	BasicNPC NPCType = iota
	// Add more NPC types here later
)

type MovementState int

const (
	Idle MovementState = iota
	Wandering
	Pausing
	Pursuing
)

type NPC struct {
	BaseEntity
	Type           NPCType
	DetectionRange float64
	State          MovementState
	WanderAngle    float64             // Current wandering direction
	StateTimer     float64             // Renamed from WanderTimer for clarity
	LastUpdate     time.Time           // For time-based updates
	StartPos       properties.Position // Store starting position for wandering radius check
	Mass           float64             // Added for push mechanics
}

const (
	WanderSpeed         = 1.0   // Slower than pursuit speed
	WanderRadius        = 100.0 // How far they wander from their start
	MinWanderTime       = 2.0   // Minimum time before changing direction
	MaxWanderTime       = 5.0   // Maximum time before changing direction
	MinPauseTime        = 1.0   // Minimum time to pause
	MaxPauseTime        = 3.0   // Maximum time to pause
	NPCAvoidRadius      = 25.0  // How far NPCs try to stay from each other
	NPCAvoidStrength    = 0.5   // How strongly NPCs avoid each other
	PlayerAvoidRadius   = 20.0  // How far NPCs try to stay from player
	PlayerAvoidStrength = 1.0   // How strongly NPCs avoid player (stronger than NPC avoidance)
)

func NewNPC(npcType NPCType, x, y float64) *NPC {
	return &NPC{
		BaseEntity: BaseEntity{
			Position: properties.Position{X: x, Y: y},
			Size:     16.0,
		},
		Type:           npcType,
		DetectionRange: 150.0,
		State:          Idle,
		WanderAngle:    rand.Float64() * math.Pi * 2,
		StateTimer:     rand.Float64()*(MaxWanderTime-MinWanderTime) + MinWanderTime,
		LastUpdate:     time.Now(),
		StartPos:       properties.Position{X: x, Y: y},
		Mass:           1.0, // Default mass, can be varied for different NPC types
	}
}

func (n *NPC) resolveCollisions(otherNPCs []*NPC) {
	for _, other := range otherNPCs {
		if other == n {
			continue
		}

		// Calculate distance between NPCs
		dx := n.Position.X - other.Position.X
		dy := n.Position.Y - other.Position.Y
		dist := math.Sqrt(dx*dx + dy*dy)

		// Check for collision
		minDist := (n.Size + other.Size) / 2
		if dist < minDist && dist > 0 {
			// Calculate overlap
			overlap := minDist - dist

			// Calculate normalized direction
			nx := dx / dist
			ny := dy / dist

			// Calculate mass ratio for push distribution
			totalMass := n.Mass + other.Mass
			ratioSelf := other.Mass / totalMass
			ratioOther := n.Mass / totalMass

			// Push both NPCs apart
			pushX := nx * overlap * 0.5 // Half the overlap for each NPC
			pushY := ny * overlap * 0.5

			// Apply push based on mass ratio
			n.Position.X += pushX * ratioSelf
			n.Position.Y += pushY * ratioSelf
			other.Position.X -= pushX * ratioOther
			other.Position.Y -= pushY * ratioOther

			// Adjust velocities to prevent immediate re-collision
			// Transfer some momentum
			dotProduct := (n.VelX*nx + n.VelY*ny) - (other.VelX*nx + other.VelY*ny)
			if dotProduct > 0 {
				impulse := dotProduct * 0.5 // Bounce factor

				// Apply impulse based on mass ratio
				n.VelX -= nx * impulse * ratioSelf
				n.VelY -= ny * impulse * ratioSelf
				other.VelX += nx * impulse * ratioOther
				other.VelY += ny * impulse * ratioOther
			}
		}
	}
}

func (n *NPC) Update(playerX, playerY float64, otherNPCs []*NPC) {
	now := time.Now()
	deltaTime := now.Sub(n.LastUpdate).Seconds()
	n.LastUpdate = now

	// First resolve any collisions
	n.resolveCollisions(otherNPCs)

	// Calculate distance to player
	distToPlayer := math.Sqrt(math.Pow(n.Position.X-playerX, 2) + math.Pow(n.Position.Y-playerY, 2))

	// Check for collision with player
	isColliding := distToPlayer <= (n.Size+PlayerSize)/2

	// If colliding with player, push NPC away slightly
	if isColliding {
		// Calculate normalized direction away from player
		dx := n.Position.X - playerX
		dy := n.Position.Y - playerY
		length := math.Sqrt(dx*dx + dy*dy)
		if length > 0 {
			// Move NPC to just outside collision radius
			pushDistance := (n.Size+PlayerSize)/2 - length
			n.Position.X += (dx / length) * pushDistance
			n.Position.Y += (dy / length) * pushDistance
		}
		// Stop NPC movement
		n.VelX = 0
		n.VelY = 0
		n.TargetVelX = 0
		n.TargetVelY = 0
		return // Skip rest of update if colliding
	}

	// Update state based on player distance
	if distToPlayer <= n.DetectionRange {
		n.State = Pursuing
	} else if distToPlayer > n.DetectionRange*1.2 {
		if n.State == Pursuing {
			n.State = Idle
		}
	}

	// Calculate avoidance vectors
	var avoidX, avoidY float64

	// Player avoidance (stronger than NPC avoidance)
	if distToPlayer < PlayerAvoidRadius {
		dx := n.Position.X - playerX
		dy := n.Position.Y - playerY
		if distToPlayer > 0 {
			strength := (PlayerAvoidRadius - distToPlayer) / PlayerAvoidRadius * PlayerAvoidStrength
			avoidX += (dx / distToPlayer) * strength
			avoidY += (dy / distToPlayer) * strength
		}
	}

	// Other NPCs avoidance
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

	switch n.State {
	case Pursuing:
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

		n.TargetVelX = dx * MaxSpeed * 0.7
		n.TargetVelY = dy * MaxSpeed * 0.7

	case Idle:
		n.StateTimer -= deltaTime
		if n.StateTimer <= 0 {
			n.State = Wandering
			n.StateTimer = rand.Float64()*(MaxWanderTime-MinWanderTime) + MinWanderTime
			n.WanderAngle = rand.Float64() * math.Pi * 2
		}
		n.TargetVelX = 0
		n.TargetVelY = 0

	case Wandering:
		n.StateTimer -= deltaTime
		if n.StateTimer <= 0 {
			n.State = Pausing
			n.StateTimer = rand.Float64()*(MaxPauseTime-MinPauseTime) + MinPauseTime
			n.TargetVelX = 0
			n.TargetVelY = 0
			break
		}

		// Check if we're too far from start position
		distFromStart := math.Sqrt(math.Pow(n.Position.X-n.StartPos.X, 2) + math.Pow(n.Position.Y-n.StartPos.Y, 2))
		if distFromStart > WanderRadius {
			// Head back toward start position
			dx := n.StartPos.X - n.Position.X
			dy := n.StartPos.Y - n.Position.Y
			length := math.Sqrt(dx*dx + dy*dy)
			if length != 0 {
				n.WanderAngle = math.Atan2(dy, dx)
			}
		} else {
			// Occasionally adjust direction slightly
			if rand.Float64() < 0.02 { // 2% chance per update
				n.WanderAngle += (rand.Float64() - 0.5) * math.Pi / 2
			}
		}

		dx := math.Cos(n.WanderAngle)
		dy := math.Sin(n.WanderAngle)

		// Add avoidance vector
		dx += avoidX
		dy += avoidY

		length := math.Sqrt(dx*dx + dy*dy)
		if length != 0 {
			dx /= length
			dy /= length
		}

		n.TargetVelX = dx * WanderSpeed
		n.TargetVelY = dy * WanderSpeed

	case Pausing:
		n.StateTimer -= deltaTime
		if n.StateTimer <= 0 {
			n.State = Wandering
			n.StateTimer = rand.Float64()*(MaxWanderTime-MinWanderTime) + MinWanderTime
			// Slightly adjust direction after pausing
			n.WanderAngle += (rand.Float64() - 0.5) * math.Pi
		}
		n.TargetVelX = 0
		n.TargetVelY = 0
	}

	n.UpdateMovement()
}

func (n *NPC) UpdateMovement() {
	nextX := n.Position.X
	nextY := n.Position.Y

	// Update X velocity
	if n.VelX < n.TargetVelX {
		n.VelX = min(n.VelX+Acceleration, n.TargetVelX)
	} else if n.VelX > n.TargetVelX {
		n.VelX = max(n.VelX-Deceleration, n.TargetVelX)
	}

	// Update Y velocity
	if n.VelY < n.TargetVelY {
		n.VelY = min(n.VelY+Acceleration, n.TargetVelY)
	} else if n.VelY > n.TargetVelY {
		n.VelY = max(n.VelY-Deceleration, n.TargetVelY)
	}

	// Apply velocity with some damping to prevent excessive bouncing
	damping := 0.98
	n.VelX *= damping
	n.VelY *= damping

	// Calculate next position
	nextX += n.VelX
	nextY += n.VelY

	// Bound checking
	nextX = max(0, min(nextX, 800-n.Size))
	nextY = max(0, min(nextY, 600-n.Size))

	// Update position
	n.Position.X = nextX
	n.Position.Y = nextY
}
