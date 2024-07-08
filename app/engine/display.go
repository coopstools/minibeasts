package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

func Display() {
	var fc uint = 0
	game := Game{imgLookup: LoadAssets(), frameCount: &fc, player: &Sprite{Imgs: []string{
		"drake_side1", "drake_side2", "drake_step1", "drake_step2"}}}
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Mini Beasts")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	player     *Sprite
	imgLookup  map[string]*ebiten.Image
	frameCount *uint
}

func (g Game) Update() error {
	var moved bool
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.MoveRight()
		moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.MoveLeft()
		moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.MoveUp()
		moved = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.MoveDown()
		moved = true
	}
	if !moved {
		g.player.ResetVelocity()
	}
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	g.drawGround(screen)
	*g.frameCount += 1
	if *g.frameCount%8 == 0 {
		g.player.UpdateLegs()
	}

	imgName, playerOpt := g.player.DisplayNameAndOpt()
	screen.DrawImage(g.imgLookup[imgName], playerOpt)
}

func (g Game) drawGround(screen *ebiten.Image) {
	img := g.imgLookup["ground_1_16x16"]
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			opt := ebiten.DrawImageOptions{}
			opt.GeoM.Scale(4.0, 4.0)
			opt.GeoM.Translate(float64(64*i), float64(64*j))
			screen.DrawImage(img, &opt)
		}
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	screenWidth, screenHeight = 500, 500
	return
}
