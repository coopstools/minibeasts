package battle

import (
	"github.com/coopstools/minibeast/app/engine/asset"
	"github.com/coopstools/minibeast/app/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	cursorPos int
	frame     int8
	pressed   *pressed
}

type pressed struct {
	key  ebiten.Key
	time int
}

func (s *State) Update(isPressed func(key ebiten.Key) bool) (string, error) {
	if s.pressed != nil && isPressed(s.pressed.key) && s.pressed.time > 0 {
		s.pressed.time -= 1
		return scene.BATTLE_STATE, nil
	}
	s.pressed = nil
	if isPressed(ebiten.KeyEscape) {
		return scene.TALL_GRASS_STATE, nil
	}
	if isPressed(ebiten.KeyW) {
		s.cursorPos += 1
		s.pressed = &pressed{
			key:  ebiten.KeyW,
			time: 20,
		}
	}
	if isPressed(ebiten.KeyS) {
		s.cursorPos -= 1
		s.pressed = &pressed{
			key:  ebiten.KeyS,
			time: 20,
		}
	}
	s.cursorPos %= 3
	return scene.BATTLE_STATE, nil
}

func (s *State) Draw(screen *ebiten.Image, imgLookup map[string]*ebiten.Image) {
	textBuilder := asset.NewBuilder(3)
	carrot := []string{">", " ", " "}
	for i, line := range []string{
		"A Line for testing!", "And another line!!!", "A line to finish things out.",
	} {
		textBuilder.Println(carrot[((i+s.cursorPos)%3+3)%3] + line)
	}
	textBuilder.Finish(screen.DrawImage)
}

func NewState() (string, State) {
	return scene.BATTLE_STATE, State{}
}
