package models

import "math/rand"

type NPCManager struct {
	NPCs []*NPC
}

func NewNPCManager() *NPCManager {
	return &NPCManager{
		NPCs: make([]*NPC, 0),
	}
}

func (m *NPCManager) SpawnNPC(npcType NPCType, x, y float64) {
	npc := NewNPC(npcType, x, y)
	m.NPCs = append(m.NPCs, npc)
}

func (m *NPCManager) SpawnRandomNPCs(count int) {
	for i := 0; i < count; i++ {
		x := rand.Float64() * (800 - PlayerSize)
		y := rand.Float64() * (600 - PlayerSize)
		m.SpawnNPC(BasicNPC, x, y)
	}
}

func (m *NPCManager) Update(playerX, playerY float64) {
	for _, npc := range m.NPCs {
		npc.Update(playerX, playerY, m.NPCs)
	}
}
