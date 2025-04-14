package engine

import (
	"log"
	"os"

	"github.com/coopstools/minibeast/internal/assets/fonts"
	"github.com/coopstools/minibeast/internal/assets/images/tiles"
	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"
	"github.com/coopstools/minibeast/internal/scenes/character_creation"
	"github.com/coopstools/minibeast/internal/scenes/profession"
	"github.com/coopstools/minibeast/internal/scenes/region"
	"github.com/coopstools/minibeast/internal/scenes/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	state        *models.GameState
	currentScene scenes.Scene
	scenes       map[string]scenes.Scene
	assets       *scenes.SceneAssets
}

func NewGame() *Game {
	tiles, err := tiles.Load()
	if err != nil {
		print("failed to load tiles: %v", err)
		os.Exit(1)
	}

	assets := &scenes.SceneAssets{
		Font:  fonts.LoadFontFace(),
		Tiles: tiles,
	}

	g := &Game{
		state:  models.NewGameState(),
		scenes: make(map[string]scenes.Scene),
		assets: assets,
	}

	// Initialize all scenes
	g.scenes["character_creation"] = character_creation.New(g.state, g)
	g.scenes["profession"] = profession.New(g.state, g)
	g.scenes["world"] = world.New(g.state, g, assets)
	g.scenes["region"] = region.New(g.state, g)
	// Set initial scene
	g.currentScene = g.scenes["character_creation"]

	return g
}

func (g *Game) SwitchTo(sceneName string) {
	if g.currentScene != nil {
		g.currentScene.Unload()
	}

	if scene, exists := g.scenes[sceneName]; exists {
		g.currentScene = scene
		if err := g.currentScene.Load(); err != nil {
			log.Printf("Failed to load scene %s: %v", sceneName, err)
		}
	}
}

func (g *Game) GetAssets() *scenes.SceneAssets {
	return g.assets
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
