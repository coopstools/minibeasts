package world

import (
	"fmt"
	"math"
	"os"

	"github.com/coopstools/minibeast/internal/assets/images/tiles"
	"github.com/coopstools/minibeast/internal/models"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene struct {
	worldData  *models.WorldData
	tileImages map[models.TileType]*ebiten.Image
}

func NewScene() *Scene {
	// Calculate world size based on screen dimensions
	worldWidth := 800 / 16  // Assuming 800px screen width
	worldHeight := 600 / 16 // Assuming 600px screen height

	scene := &Scene{
		worldData:  models.NewWorldData(worldWidth, worldHeight),
		tileImages: make(map[models.TileType]*ebiten.Image),
	}

	// Load tile images
	scene.loadTileImages()

	return scene
}

func (s *Scene) loadTileImages() {
	tiles, err := tiles.Load()
	if err != nil {
		fmt.Println("failed to load tiles for world:", err)
		os.Exit(2)
	}

	for tileType, tileImage := range tiles {
		s.tileImages[models.TileType(tileType)] = ebiten.NewImageFromImage(tileImage)
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
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
			tileImage := s.tileImages[tile.Type]
			screen.DrawImage(tileImage, opt)
		}
	}
}

func (s *Scene) Update() error {
	return nil
}
