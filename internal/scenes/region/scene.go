package region

import (
	"image/color"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	gameState    *models.GameState
	sceneManager scenes.SceneManager
}

func NewScene(gameState *models.GameState, manager scenes.SceneManager) *Scene {
	return &Scene{
		gameState:    gameState,
		sceneManager: manager,
	}
}

func (s *Scene) Update() error {
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	// For now, just draw a blank screen with different background color
	screen.Fill(color.RGBA{20, 20, 40, 255})
}
