package engine

import (
	"os"

	"github.com/coopstools/minibeast/internal/assets/fonts"
	"github.com/coopstools/minibeast/internal/assets/images/tiles"
	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/coopstools/minibeast/internal/scenes/character_creation"
	"github.com/coopstools/minibeast/internal/scenes/profession"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state        *models.GameState
	currentScene scenes.Scene
	scenes       map[string]scenes.Scene
	assets       *scenes.SceneAssets
}

func NewGame() *Game {
	tiles, err := tiles.LoadTiles()
	if err != nil {
		print("failed to load tiles: %v", err)
		os.Exit(1)
	}

	g := &Game{
		state:  models.NewGameState(),
		scenes: make(map[string]scenes.Scene),
		assets: &scenes.SceneAssets{
			Font:  fonts.LoadFontFace(),
			Tiles: tiles,
		},
	}

	// Initialize all scenes
	g.scenes["character_creation"] = character_creation.NewScene(g.state, g)
	g.scenes["profession"] = profession.NewScene(g.state, g)

	// Set initial scene
	g.currentScene = g.scenes["character_creation"]

	return g
}

func (g *Game) SwitchTo(sceneName string) {
	if scene, exists := g.scenes[sceneName]; exists {
		g.currentScene = scene
	}
}

func (g *Game) Update() error {
	return g.currentScene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentScene.Draw(screen, g.assets)
}

func (g *Game) Layout(w, h int) (int, int) {
	return 800, 600
}
