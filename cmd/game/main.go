package main

import (
	"log"

	"github.com/coopstools/minibeast/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := engine.NewGame()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Character Creator")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
