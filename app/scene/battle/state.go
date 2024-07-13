package battle

import (
	"github.com/coopstools/minibeast/app/engine/asset"
	"github.com/coopstools/minibeast/app/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct{}

func (s State) Update(isPressed func(key ebiten.Key) bool) (string, error) {
	if isPressed(ebiten.KeyEscape) {
		return scene.TALL_GRASS_STATE, nil
	}
	return scene.BATTLE_STATE, nil
}

func (s State) Draw(screen *ebiten.Image, imgLookup map[string]*ebiten.Image) {
	textBuilder := asset.NewBuilder(3)
	textBuilder.Println(">A Line for testing!")
	textBuilder.Println(" And another line!!!")
	textBuilder.Println(" A line to finish things out.")
	textBuilder.Finish(screen.DrawImage)
}

func NewState() (string, State) {
	return scene.BATTLE_STATE, State{}
}
