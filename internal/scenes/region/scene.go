package region

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/coopstools/minibeast/internal/entities/npcs/hostile"
	"github.com/coopstools/minibeast/internal/entities/player"
	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/coopstools/minibeast/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Scene struct {
	gameState     *models.GameState
	sceneManager  scenes.SceneManager
	player        *player.Player
	npcManager    *hostile.Manager
	statsPopup    *ui.Popup
	healthDisplay *ui.HealthDisplay
	isLoaded      bool
}

func New(gameState *models.GameState, manager scenes.SceneManager) *Scene {
	return &Scene{
		gameState:    gameState,
		sceneManager: manager,
	}
}

func (s *Scene) Load() error {
	if s.isLoaded {
		return nil
	}

	// Create player
	s.player = player.New(s.gameState.Character)

	// Create NPC manager
	s.npcManager = hostile.NewManager()
	s.npcManager.SpawnRandomNPCs(5)

	// Create UI elements
	assets := s.sceneManager.GetAssets()
	s.statsPopup = ui.NewPopup(400, 300, assets)
	s.healthDisplay = ui.NewHealthDisplay(10, 20, assets.Font)

	s.isLoaded = true
	return nil
}

func (s *Scene) Unload() error {
	s.isLoaded = false
	return nil
}

func (s *Scene) Update() error {
	if !s.isLoaded {
		return nil
	}

	// Handle stats popup
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		if !s.statsPopup.Visible {
			stats := s.gameState.Character
			content := fmt.Sprintf(
				"Character Stats\n"+
					"Name: %s\n"+
					"Profession: %s\n"+
					"Strength: %d\n"+
					"Dexterity: %d\n"+
					"Vitality: %d\n"+
					"Intelligence: %d",
				stats.Name,
				stats.Profession,
				stats.Strength,
				stats.Dexterity,
				stats.Vitality,
				stats.Intelligence,
			)
			s.statsPopup.Show(200, 150, content)
		} else {
			s.statsPopup.Hide()
		}
	}

	// Update player movement and knockback
	if s.player.Knockback.Active {
		now := time.Now()
		if now.Sub(s.player.Knockback.StartTime) >= s.player.Knockback.Duration {
			s.player.Knockback.Active = false
		} else {
			// Apply knockback movement
			s.player.Position.X += s.player.Knockback.VelX
			s.player.Position.Y += s.player.Knockback.VelY

			// Gradually reduce knockback velocity
			reduction := 1.0 - (now.Sub(s.player.Knockback.StartTime).Seconds() / s.player.Knockback.Duration.Seconds())
			s.player.Knockback.VelX *= reduction
			s.player.Knockback.VelY *= reduction

			// Bound checking during knockback
			s.player.Position.X = max(0, min(s.player.Position.X, 800-s.player.Size))
			s.player.Position.Y = max(0, min(s.player.Position.Y, 600-s.player.Size))
		}
	}

	// Only update normal movement if not in knockback
	if !s.player.Knockback.Active {
		if err := s.player.Update(); err != nil {
			return err
		}
	}

	// Update NPCs
	if err := s.npcManager.Update(s.player.Position); err != nil {
		return err
	}

	// Handle collisions
	s.handleCollisions()

	return nil
}

func (s *Scene) handleCollisions() {
	// Check collisions between player and NPCs
	for _, npc := range s.npcManager.NPCs {
		dx := s.player.Position.X - npc.Position.X
		dy := s.player.Position.Y - npc.Position.Y
		dist := math.Sqrt(dx*dx + dy*dy)
		minDist := (s.player.Size + npc.Size) / 2

		// If NPC is touching player
		if dist < minDist {
			nx := dx / dist
			ny := dy / dist
			// First resolve the collision by pushing entities apart
			if dist > 0 {
				overlap := minDist - dist

				// Push player and NPC apart (player gets pushed more since NPCs are "heavier")
				pushRatio := 0.7 // Player takes 70% of the push
				s.player.Position.X += nx * overlap * pushRatio
				s.player.Position.Y += ny * overlap * pushRatio
				npc.Position.X -= nx * overlap * (1 - pushRatio)
				npc.Position.Y -= ny * overlap * (1 - pushRatio)
			}

			// Then handle damage and knockback if player can take damage
			if s.player.Health.CanTakeDamage() {
				// Apply damage
				s.player.Health.TakeDamage(npc.DamageAmount)

				// Apply knockback
				s.player.Knockback = player.KnockbackState{
					Active:    true,
					VelX:      nx * 8.0, // Knockback force
					VelY:      ny * 8.0,
					Duration:  time.Second / 2, // 0.5 second knockback
					StartTime: time.Now(),
				}
			}
		}
	}

	// Ensure player stays within screen bounds
	s.player.Position.X = max(0, min(s.player.Position.X, 800-s.player.Size))
	s.player.Position.Y = max(0, min(s.player.Position.Y, 600-s.player.Size))
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

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	// Draw background
	screen.Fill(color.RGBA{20, 20, 40, 255})

	// Draw NPCs
	for _, npc := range s.npcManager.NPCs {
		npc.Draw(screen)
	}

	// Draw player with knockback effect
	playerColor := color.RGBA{255, 0, 0, 255}
	if s.player.Knockback.Active {
		// Flash white when in knockback
		playerColor = color.RGBA{255, 255, 255, 255}
	}

	vector.DrawFilledCircle(screen,
		float32(s.player.Position.X+s.player.Size/2),
		float32(s.player.Position.Y+s.player.Size/2),
		float32(s.player.Size/2),
		playerColor,
		false)

	// Draw health display
	s.healthDisplay.Draw(screen,
		s.gameState.Character.Health.Current,
		s.gameState.Character.Health.Max)

	// Draw stats popup if visible
	if s.statsPopup.Visible {
		s.statsPopup.Draw(screen, assets)
	}
}
