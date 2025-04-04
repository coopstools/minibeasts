package scenes

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type SceneManager interface {
	SwitchTo(sceneName string)
}

type SceneAssets struct {
	Font  font.Face
	Tiles map[string]image.Image
}

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image, assets *SceneAssets)
}
