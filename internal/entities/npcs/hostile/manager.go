package hostile

import (
	"math"
	"math/rand"

	"github.com/coopstools/minibeast/internal/models/properties"
)

type Manager struct {
	NPCs []*NPC
}

func NewManager() *Manager {
	return &Manager{
		NPCs: make([]*NPC, 0),
	}
}

func (m *Manager) Update(playerPos properties.Position) error {
	// First update all NPCs' behavior
	for _, npc := range m.NPCs {
		if err := npc.Update(playerPos.X, playerPos.Y, m.NPCs); err != nil {
			return err
		}
	}

	// Then resolve collisions between NPCs
	m.resolveCollisions()

	return nil
}

func (m *Manager) resolveCollisions() {
	// Check each pair of NPCs for collisions
	for i := 0; i < len(m.NPCs); i++ {
		for j := i + 1; j < len(m.NPCs); j++ {
			npc1 := m.NPCs[i]
			npc2 := m.NPCs[j]

			dx := npc1.Position.X - npc2.Position.X
			dy := npc1.Position.Y - npc2.Position.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			minDist := (npc1.Size + npc2.Size) / 2

			if dist < minDist && dist > 0 {
				// Calculate overlap
				overlap := minDist - dist

				// Calculate normalized direction
				nx := dx / dist
				ny := dy / dist

				// Push both NPCs apart
				pushX := nx * overlap * 0.5 // Half the overlap for each NPC
				pushY := ny * overlap * 0.5

				// Apply the push
				npc1.Position.X += pushX
				npc1.Position.Y += pushY
				npc2.Position.X -= pushX
				npc2.Position.Y -= pushY

				// Adjust velocities to prevent immediate re-collision
				dotProduct := (npc1.VelX*nx + npc1.VelY*ny) - (npc2.VelX*nx + npc2.VelY*ny)
				if dotProduct > 0 {
					impulse := dotProduct * 0.5 // Bounce factor
					npc1.VelX -= nx * impulse
					npc1.VelY -= ny * impulse
					npc2.VelX += nx * impulse
					npc2.VelY += ny * impulse
				}
			}
		}
	}
}

func (m *Manager) SpawnRandomNPCs(count int) {
	for i := 0; i < count; i++ {
		x := rand.Float64() * (800 - 16) // Screen width - NPC size
		y := rand.Float64() * (600 - 16) // Screen height - NPC size
		npc := New(Basic, x, y)
		m.NPCs = append(m.NPCs, npc)
	}
}
