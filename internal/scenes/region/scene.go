package region

import (
	"fmt"
	"image/color"
	"math"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/models/properties"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/coopstools/minibeast/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Scene struct {
	gameState     *models.GameState
	sceneManager  scenes.SceneManager
	playerPos     models.PlayerPosition
	statsPopup    *ui.Popup
	npcManager    *models.NPCManager
	healthDisplay *ui.HealthDisplay
}

func NewScene(gameState *models.GameState, manager scenes.SceneManager, assets *scenes.SceneAssets) *Scene {
	scene := &Scene{
		gameState:    gameState,
		sceneManager: manager,
		playerPos: models.PlayerPosition{
			Position: properties.Position{X: 400, Y: 300}, // Start in middle of screen
		},
		npcManager:    models.NewNPCManager(),
		healthDisplay: ui.NewHealthDisplay(10, 20, assets.Font),
	}

	// Create a larger popup for character stats
	scene.statsPopup = ui.NewPopup(400, 300, assets)

	// Spawn some initial NPCs
	scene.npcManager.SpawnRandomNPCs(5)

	return scene
}

func (s *Scene) Update() error {
	// Reset target velocities
	s.playerPos.TargetVelX = 0
	s.playerPos.TargetVelY = 0

	// Handle diagonal movement normalization
	inputX := 0.0
	inputY := 0.0

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		inputY--
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		inputY++
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		inputX--
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		inputX++
	}

	// Normalize diagonal movement
	if inputX != 0 && inputY != 0 {
		length := math.Sqrt(inputX*inputX + inputY*inputY)
		inputX /= length
		inputY /= length
	}

	// Set target velocities
	s.playerPos.TargetVelX = inputX * models.MaxSpeed
	s.playerPos.TargetVelY = inputY * models.MaxSpeed

	// Update movement
	s.playerPos.UpdateMovement()

	// Update NPCs
	s.npcManager.Update(s.playerPos.Position.X, s.playerPos.Position.Y)

	// Check for NPC collisions and damage
	for _, npc := range s.npcManager.NPCs {
		dx := s.playerPos.Position.X - npc.Position.X
		dy := s.playerPos.Position.Y - npc.Position.Y
		dist := math.Sqrt(dx*dx + dy*dy)

		if dist <= (models.PlayerSize+npc.Size)/2 {
			if s.gameState.Character.Health.CanTakeDamage() {
				s.gameState.Character.Health.TakeDamage(npc.DamageAmount)
				// Apply knockback when taking damage
				s.playerPos.ApplyKnockback(npc.Position.X, npc.Position.Y)
			}
		}
	}

	// Toggle stats popup with E key
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		if !s.statsPopup.Visible {
			content := fmt.Sprintf(
				"Character Information\n\n"+
					"Name: %s\n"+
					"Profession: %s\n\n"+
					"Stats:\n"+
					"Strength: %d\n"+
					"Dexterity: %d\n"+
					"Vitality: %d\n"+
					"Intelligence: %d",
				s.gameState.Character.Name,
				s.gameState.Character.Profession,
				s.gameState.Character.Strength,
				s.gameState.Character.Dexterity,
				s.gameState.Character.Vitality,
				s.gameState.Character.Intelligence,
			)
			s.statsPopup.Show(200, 150, content) // Center on screen
		} else {
			s.statsPopup.Hide()
		}
	}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	// Draw background
	screen.Fill(color.RGBA{20, 20, 40, 255})

	// Draw health display
	s.healthDisplay.Draw(screen,
		s.gameState.Character.Health.Current,
		s.gameState.Character.Health.Max)

	// Draw NPCs
	for _, npc := range s.npcManager.NPCs {
		var npcColor color.RGBA

		// Check for collisions with other NPCs
		isColliding := false
		for _, other := range s.npcManager.NPCs {
			if other == npc {
				continue
			}
			dx := npc.Position.X - other.Position.X
			dy := npc.Position.Y - other.Position.Y
			dist := math.Sqrt(dx*dx + dy*dy)
			if dist <= (npc.Size+other.Size)/2 {
				isColliding = true
				break
			}
		}

		// Determine color based on state and collision
		switch npc.State {
		case models.Pursuing:
			if isColliding {
				npcColor = color.RGBA{200, 0, 200, 255} // Brightest when colliding
			} else {
				npcColor = color.RGBA{160, 0, 160, 255}
			}
		case models.Wandering:
			if isColliding {
				npcColor = color.RGBA{150, 0, 150, 255}
			} else {
				npcColor = color.RGBA{128, 0, 128, 255}
			}
		case models.Pausing:
			if isColliding {
				npcColor = color.RGBA{120, 0, 120, 255}
			} else {
				npcColor = color.RGBA{100, 0, 100, 255}
			}
		case models.Idle:
			if isColliding {
				npcColor = color.RGBA{100, 0, 100, 255}
			} else {
				npcColor = color.RGBA{80, 0, 80, 255}
			}
		}

		vector.DrawFilledCircle(screen,
			float32(npc.Position.X+npc.Size/2),
			float32(npc.Position.Y+npc.Size/2),
			float32(npc.Size/2),
			npcColor,
			false)
	}

	// Draw player with knockback effect
	playerColor := color.RGBA{255, 0, 0, 255}
	if s.playerPos.Knockback.Active {
		// Flash white when in knockback
		playerColor = color.RGBA{255, 255, 255, 255}
	}

	vector.DrawFilledCircle(screen,
		float32(s.playerPos.Position.X+models.PlayerSize/2),
		float32(s.playerPos.Position.Y+models.PlayerSize/2),
		float32(models.PlayerSize/2),
		playerColor,
		false)

	// Draw stats popup if visible
	if s.statsPopup.Visible {
		s.statsPopup.Draw(screen, assets)
	}
}
