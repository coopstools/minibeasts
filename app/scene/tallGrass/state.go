package tallGrass

import (
	"github.com/coopstools/minibeast/app/scene"
	"github.com/coopstools/minibeast/app/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type State struct {
	player     *scene.Sprite
	frameCount *uint

	sharedState *state.Shared
}

func (s State) Update(isPressed func(key ebiten.Key) bool) (string, error) {
	if isPressed(ebiten.KeyE) {
		return scene.BATTLE_STATE, nil
	}

	var moved bool
	if isPressed(ebiten.KeyD) {
		s.player.MoveRight()
		moved = true
	}
	if isPressed(ebiten.KeyA) {
		s.player.MoveLeft()
		moved = true
	}
	if isPressed(ebiten.KeyW) {
		s.player.MoveUp()
		moved = true
	}
	if isPressed(ebiten.KeyS) {
		s.player.MoveDown()
		moved = true
	}
	if isPressed(ebiten.KeyE) {

	}
	if !moved {
		s.player.ResetVelocity()
	}
	return scene.TALL_GRASS_STATE, nil
}

func (s State) Draw(screen *ebiten.Image) {
	s.drawGround(screen, s.sharedState.BackGroundImages["ground_1_16x16"])
	*s.frameCount += 1
	if *s.frameCount%8 == 0 {
		s.player.UpdateLegs()
	}

	imgName, playerOpt := s.player.DisplayNameAndOpt()
	screen.DrawImage(s.sharedState.PullSelectedCharacter(imgName), playerOpt)
}

func (s State) drawGround(screen *ebiten.Image, img *ebiten.Image) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			opt := ebiten.DrawImageOptions{}
			opt.GeoM.Scale(4.0, 4.0)
			opt.GeoM.Translate(float64(64*i), float64(64*j))
			screen.DrawImage(img, &opt)
		}
	}
}

func NewState(s *state.Shared, player *scene.Sprite) (string, State) {
	var fc uint
	return scene.TALL_GRASS_STATE, State{
		player:      player,
		frameCount:  &fc,
		sharedState: s,
	}
}
