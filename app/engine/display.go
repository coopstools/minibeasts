package engine

import (
	"github.com/coopstools/minibeast/app/engine/asset"
	"github.com/coopstools/minibeast/app/scene"
	"github.com/coopstools/minibeast/app/scene/battle"
	"github.com/coopstools/minibeast/app/scene/tallGrass"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

func Display() {
	states := make(map[string]state, 2)
	tgName, tallGrassScene := tallGrass.NewState(&scene.Sprite{Imgs: []string{
		"drake_side1", "drake_side2", "drake_step1", "drake_step2",
	}})
	states[tgName] = tallGrassScene
	bName, battleScene := battle.NewState()
	states[bName] = battleScene
	curState := tgName
	game := Game{
		curState:  &curState,
		imgLookup: asset.LoadAssets(),
		states:    states,
	}
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Mini Beasts")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type state interface {
	Update(isPressed func(key ebiten.Key) bool) (string, error)
	Draw(screen *ebiten.Image, imgLookup map[string]*ebiten.Image)
}

type Game struct {
	states    map[string]state
	curState  *string
	imgLookup map[string]*ebiten.Image
}

func (g Game) Update() error {
	newState, err := g.states[*g.curState].Update(ebiten.IsKeyPressed)
	*g.curState = newState
	return err
}

func (g Game) Draw(screen *ebiten.Image) {
	g.states[*g.curState].Draw(screen, g.imgLookup)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = 500, 500
	return
}
