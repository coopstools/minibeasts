package battle

import (
	"github.com/coopstools/minibeast/app/asset"
	"github.com/coopstools/minibeast/app/scene"
	"github.com/coopstools/minibeast/app/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	mvSelPopup Popup2x2

	cursorPos int
	frame     int8
	pressed   *pressed

	sharedState *state.Shared
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
		s.sharedState.SetSelected(s.mvSelPopup.PullCurrent())
		return scene.TALL_GRASS_STATE, nil
	}
	for _, k := range []struct {
		mv  func()
		key ebiten.Key
	}{
		{mv: s.mvSelPopup.MoveUp, key: ebiten.KeyW},
		{mv: s.mvSelPopup.MoveUp, key: ebiten.KeyS},
		{mv: s.mvSelPopup.MoveLeft, key: ebiten.KeyA},
		{mv: s.mvSelPopup.MoveLeft, key: ebiten.KeyD},
	} {
		if !isPressed(k.key) {
			continue
		}
		k.mv()
		s.pressed = &pressed{
			key:  k.key,
			time: 20,
		}
	}
	return scene.BATTLE_STATE, nil
}

func (s *State) Draw(screen *ebiten.Image) {
	selImg := s.mvSelPopup.Draw()
	screen.DrawImage(selImg, &ebiten.DrawImageOptions{})
}

func NewState(s *state.Shared) (string, State) {
	// popup := New2x2Popup(asset.LoadFont(), []string{"Tail Whip", "Growl", "Tackl", "Quick Strike"}, 200, 200)
	popup := New2x2Popup(asset.LoadFont(), s.PullAvailable(), 300, 200)
	return scene.BATTLE_STATE, State{mvSelPopup: popup, sharedState: s}
}
