package world

import (
	"math"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	worldData    *models.WorldData
	gameState    *models.GameState
	sceneManager scenes.SceneManager
}

func NewScene(gameState *models.GameState, sceneManager scenes.SceneManager) *Scene {
	// Calculate world size based on screen dimensions
	worldWidth := 800 / 16  // Assuming 800px screen width
	worldHeight := 600 / 16 // Assuming 600px screen height

	scene := &Scene{
		worldData:    models.NewWorldData(worldWidth, worldHeight),
		gameState:    gameState,
		sceneManager: sceneManager,
	}

	return scene
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	// Create options for drawing tiles
	opt := &ebiten.DrawImageOptions{}

	// Draw each tile
	for y, row := range s.worldData.Tiles {
		for x, tile := range row {
			opt.GeoM.Reset()

			// Set position
			opt.GeoM.Translate(float64(x*16), float64(y*16))

			// Set rotation around center
			if tile.Rotation != 0 {
				opt.GeoM.Translate(-8, -8) // Move to center
				opt.GeoM.Rotate(float64(tile.Rotation) * math.Pi / 180)
				opt.GeoM.Translate(8, 8) // Move back
			}

			// Draw the tile
			tileImage := assets.Tiles[string(tile.Type)]
			//TODO: convert map of images to map of ebiten images
			ebitImage := ebiten.NewImageFromImage(tileImage)
			screen.DrawImage(ebitImage, opt)
		}
	}
}

func (s *Scene) Update() error {
	return nil
}
