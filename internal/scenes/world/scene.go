package world

import (
	"image/color"
	"math"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Scene struct {
	worldData    *models.WorldData
	gameState    *models.GameState
	sceneManager scenes.SceneManager
	hoveredTile  struct {
		X, Y int
	}
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

			// Set rotation around center
			if tile.Rotation != 0 {
				opt.GeoM.Translate(-8, -8) // Move to center
				opt.GeoM.Rotate(float64(tile.Rotation) * math.Pi / 180)
				opt.GeoM.Translate(8, 8) // Move back
			}

			// Set position
			opt.GeoM.Translate(float64(x*16), float64(y*16))

			// Draw the tile
			tileImage := assets.Tiles[string(tile.Type)]
			//TODO: convert map of images to map of ebiten images
			ebitImage := ebiten.NewImageFromImage(tileImage)
			screen.DrawImage(ebitImage, opt)
		}
	}

	// Draw highlight for hovered tile
	if s.isValidTilePosition(s.hoveredTile.X, s.hoveredTile.Y) {
		s.drawTileHighlight(screen, s.hoveredTile.X, s.hoveredTile.Y, color.RGBA{255, 0, 0, 255})
	}
}

func (s *Scene) isValidTilePosition(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(s.worldData.Tiles) && x < len(s.worldData.Tiles[0])
}

func (s *Scene) drawTileHighlight(screen *ebiten.Image, tileX, tileY int, highlightColor color.Color) {
	x := float32(tileX * 16)
	y := float32(tileY * 16)

	// Draw four lines to create a border
	vector.StrokeLine(screen, x, y, x+16, y, 1, highlightColor, false)       // Top
	vector.StrokeLine(screen, x+16, y, x+16, y+16, 1, highlightColor, false) // Right
	vector.StrokeLine(screen, x, y+16, x+16, y+16, 1, highlightColor, false) // Bottom
	vector.StrokeLine(screen, x, y, x, y+16, 1, highlightColor, false)       // Left
}

func (s *Scene) Update() error {
	// Get current mouse position
	x, y := ebiten.CursorPosition()

	// Convert mouse position to tile coordinates
	s.hoveredTile.X = x / 16 // Since tiles are 16x16
	s.hoveredTile.Y = y / 16

	// Ensure coordinates are within bounds
	if s.hoveredTile.X >= len(s.worldData.Tiles[0]) {
		s.hoveredTile.X = len(s.worldData.Tiles[0]) - 1
	}
	if s.hoveredTile.Y >= len(s.worldData.Tiles) {
		s.hoveredTile.Y = len(s.worldData.Tiles) - 1
	}

	return nil
}
