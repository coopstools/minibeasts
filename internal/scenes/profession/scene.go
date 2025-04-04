package profession

import (
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
	professions  []string
}

func NewScene(state *models.GameState, manager scenes.SceneManager) *Scene {
	return &Scene{
		gameState:    state,
		sceneManager: manager,
		selected:     0,
		keyDelay:     time.Millisecond * 200,
		professions:  []string{"Farmer", "Miner", "Blacksmith", "Merchant"},
	}
}

func (s *Scene) Update() error {
	now := time.Now()
	if now.Sub(s.lastKeyPress) >= s.keyDelay {
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			s.selected--
			if s.selected < 0 {
				s.selected = len(s.professions) - 1
			}
			s.lastKeyPress = now
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			s.selected++
			if s.selected >= len(s.professions) {
				s.selected = 0
			}
			s.lastKeyPress = now
		}
	}

	// Handle profession selection
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.gameState.Character.Profession = s.professions[s.selected]
		s.sceneManager.SwitchTo("world")
	}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {

	screen.Fill(color.RGBA{40, 40, 40, 255})

	text.Draw(screen, "Choose Your Profession", assets.Font, 100, 50, color.White)

	for i, prof := range s.professions {
		y := 100 + i*30
		if i == s.selected {
			text.Draw(screen, "> "+prof, assets.Font, 100, y, color.White)
		} else {
			text.Draw(screen, "  "+prof, assets.Font, 100, y, color.White)
		}
	}

	text.Draw(screen, "Press ENTER to confirm", assets.Font, 100, 280, color.RGBA{180, 180, 180, 255})
}
