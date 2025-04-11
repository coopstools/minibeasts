package world

import (
	"fmt"
	"image/color"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/coopstools/minibeast/internal/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Scene struct {
	worldData    *models.WorldData
	gameState    *models.GameState
	sceneManager scenes.SceneManager
	hoveredTile  struct {
		X, Y int
	}
	selectedTile struct {
		X, Y int
	}
	popup *ui.Popup
}

func NewScene(gameState *models.GameState, sceneManager scenes.SceneManager, assets *scenes.SceneAssets) *Scene {
	worldWidth := 800 / 16
	worldHeight := 600 / 16

	scene := &Scene{
		worldData:    models.NewWorldData(worldWidth, worldHeight, assets),
		gameState:    gameState,
		sceneManager: sceneManager,
		popup:        ui.NewPopup(300, 100, assets),
	}

	return scene
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	// Create options for drawing tiles
	opt := &ebiten.DrawImageOptions{}
	for y, row := range s.worldData.Tiles {
		for x, tile := range row {
			opt.GeoM.Reset()
			// TODO: preset tile position (if possible)
			opt.GeoM.Translate(float64(x*16), float64(y*16))

			tileImage := assets.Tiles[string(tile.Type)][tile.Rotation]
			screen.DrawImage(tileImage, opt)
		}
	}

	if s.popup.Visible {
		s.drawTileHighlight(screen, s.selectedTile.X, s.selectedTile.Y, color.RGBA{0, 0, 0, 255})
	}
	if s.isValidTilePosition(s.hoveredTile.X, s.hoveredTile.Y) {
		s.drawTileHighlight(screen, s.hoveredTile.X, s.hoveredTile.Y, color.RGBA{255, 0, 0, 255})
	}

	s.popup.Draw(screen, assets)
}

func (s *Scene) isValidTilePosition(x, y int) bool {
	return x >= 0 && y >= 0 && y < len(s.worldData.Tiles) && x < len(s.worldData.Tiles[0])
}

func (s *Scene) drawTileHighlight(screen *ebiten.Image, tileX, tileY int, highlightColor color.Color) {
	x := float32(tileX * 16)
	y := float32(tileY * 16)

	// Draw four lines to create a border
	vector.StrokeLine(screen, x, y, x+16, y, 3, highlightColor, false)       // Top
	vector.StrokeLine(screen, x+16, y, x+16, y+16, 3, highlightColor, false) // Right
	vector.StrokeLine(screen, x, y+16, x+16, y+16, 3, highlightColor, false) // Bottom
	vector.StrokeLine(screen, x, y, x, y+16, 3, highlightColor, false)       // Left
}

func (s *Scene) Update() error {
	// Get current mouse position
	x, y := ebiten.CursorPosition()

	newX, newY := x/16, y/16
	s.hoveredTile.X = newX
	s.hoveredTile.Y = newY

	// Ensure coordinates are within bounds
	if s.hoveredTile.X >= len(s.worldData.Tiles[0]) {
		s.hoveredTile.X = len(s.worldData.Tiles[0]) - 1
	}
	if s.hoveredTile.Y >= len(s.worldData.Tiles) {
		s.hoveredTile.Y = len(s.worldData.Tiles) - 1
	}

	if s.popup.Visible {
		// Update button hover state
		s.popup.EnterButton.Hovered = s.popup.EnterButton.Contains(x, y)

		// Handle button click
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if s.popup.EnterButton.Contains(x, y) {
				s.sceneManager.SwitchTo("region")
				return nil
			}
		}
	}

	// Handle mouse click for popup
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if s.isValidTilePosition(newX, newY) {
			s.selectedTile = s.hoveredTile

			tile := s.worldData.Tiles[newY][newX]
			content := fmt.Sprintf("Tile Position: (%d, %d)\nType: %s\nRotation: %d°",
				newX, newY, tile.Type, tile.Rotation*90)

			// Position popup near but not under the mouse
			popupX := x + 10
			popupY := y + 10

			// Keep popup on screen
			if popupX+s.popup.Width > 800 {
				popupX = 800 - s.popup.Width
			}
			if popupY+s.popup.Height > 600 {
				popupY = 600 - s.popup.Height
			}

			s.popup.Show(popupX, popupY, content)
		}
	}

	// Hide popup when right-clicking
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.popup.Hide()
	}

	return nil
}
