package engine

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed asset/*
var assetDir embed.FS

func LoadAssets() map[string]*ebiten.Image {
	fmt.Println("Loading Assets")
	dir, err := assetDir.ReadDir("asset")
	if err != nil {
		panic(err)
	}
	assets := make(map[string]*ebiten.Image, len(dir))
	for _, f := range dir {
		nl := len(f.Name())
		if nl <= 4 {
			continue
		}
		if f.Name()[nl-4:] != ".png" {
			continue
		}
		assetName := f.Name()[:nl-4]
		rawImg, err := assetDir.Open("asset/" + f.Name())
		if err != nil {
			panic(err)
		}
		ebitImg, _, err := ebitenutil.NewImageFromReader(rawImg)
		if err != nil {
			panic(err)
		}
		assets[assetName] = ebitImg
	}
	fmt.Println("Finished Loading Assets")
	return assets
}
