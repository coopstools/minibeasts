package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type SceneManager interface {
	SwitchTo(sceneName string)
	GetAssets() *SceneAssets
}

type SceneAssets struct {
	Font  font.Face
	Tiles map[string][]*ebiten.Image
}

type Scene interface {
	Load() error
	Unload() error
	Update() error
	Draw(screen *ebiten.Image, assets *SceneAssets)
}
