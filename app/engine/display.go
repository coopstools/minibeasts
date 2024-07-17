package engine

import (
	"github.com/coopstools/minibeast/app/scene"
	"github.com/coopstools/minibeast/app/scene/battle"
	"github.com/coopstools/minibeast/app/scene/tallGrass"
	game_state "github.com/coopstools/minibeast/app/state"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

func Display() {

	sharedState := game_state.NewShared()

	scenes := make(map[string]state, 2)
	tgName, tallGrassScene := tallGrass.NewState(&sharedState, &scene.Sprite{Imgs: []string{
		"side1", "side2", "up", "down",
		//"drake_side1", "drake_side2", "drake_step1", "drake_step2",
	}})
	scenes[tgName] = tallGrassScene
	bName, battleScene := battle.NewState(&sharedState)
	scenes[bName] = &battleScene
	curState := tgName

	game := Game{
		curState: &curState,
		states:   scenes,
	}
	ebiten.SetWindowSize(1200, 800)
	ebiten.SetWindowTitle("Mini Beasts")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type state interface {
	Update(isPressed func(key ebiten.Key) bool) (string, error)
	Draw(screen *ebiten.Image)
}

type Game struct {
	states   map[string]state
	curState *string
}

func (g Game) Update() error {
	newState, err := g.states[*g.curState].Update(ebiten.IsKeyPressed)
	*g.curState = newState
	return err
}

func (g Game) Draw(screen *ebiten.Image) {
	g.states[*g.curState].Draw(screen)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = 1200, 800
	return
}

func combine[t any](m0, m1 map[string]t) map[string]t {
	newMap := make(map[string]t, len(m0)+len(m1))
	for k, v := range m0 {
		newMap[k] = v
	}
	for k, v := range m1 {
		newMap[k] = v
	}
	return newMap
}
