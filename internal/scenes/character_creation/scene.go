package character_creation

import (
	"fmt"
	"image/color"
	"time"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Scene struct {
	gameState    *models.GameState
	sceneManager scenes.SceneManager
	selected     int
	lastKeyPress time.Time
	keyDelay     time.Duration
}

func NewScene(state *models.GameState, manager scenes.SceneManager) *Scene {
	return &Scene{
		gameState:    state,
		sceneManager: manager,
		selected:     0,
		keyDelay:     time.Millisecond * 200,
	}
}

func (s *Scene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.selected--
		if s.selected < 0 {
			s.selected = 2 // Number of stats - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.selected++
		if s.selected > 2 { // Number of stats - 1
			s.selected = 0
		}
	}

	// Increase/decrease selected stat
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		switch s.selected {
		case 0:
			s.gameState.Character.Strength++
		case 1:
			s.gameState.Character.Dexterity++
		case 2:
			s.gameState.Character.Vitality++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		switch s.selected {
		case 0:
			if s.gameState.Character.Strength > 1 {
				s.gameState.Character.Strength--
			}
		case 1:
			if s.gameState.Character.Dexterity > 1 {
				s.gameState.Character.Dexterity--
			}
		case 2:
			if s.gameState.Character.Vitality > 1 {
				s.gameState.Character.Vitality--
			}
		}
	}

	// Add scene transition on Enter key
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.sceneManager.SwitchTo("profession")
	}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	screen.Fill(color.RGBA{40, 40, 40, 255})

	// Draw stats
	stats := []string{
		"Strength:  " + fmt.Sprint(s.gameState.Character.Strength),
		"Dexterity: " + fmt.Sprint(s.gameState.Character.Dexterity),
		"Vitality:  " + fmt.Sprint(s.gameState.Character.Vitality),
	}

	for i, stat := range stats {
		y := 100 + i*30
		if i == s.selected {
			text.Draw(screen, "> "+stat, assets.Font, 100, y, color.White)
		} else {
			text.Draw(screen, "  "+stat, assets.Font, 100, y, color.White)
		}
	}

	// Draw instructions
	text.Draw(screen, "Use arrow keys to navigate and modify stats", assets.Font, 100, 200, color.White)

	// Add instruction for completing character creation
	text.Draw(screen, "Press ENTER when finished", assets.Font, 100, 280, color.RGBA{180, 180, 180, 255})
}
