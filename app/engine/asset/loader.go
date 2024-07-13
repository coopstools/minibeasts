package asset

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed *.png
var pngAssets embed.FS

func LoadAssets() map[string]*ebiten.Image {
	fmt.Println("Loading Assets")
	dir, err := pngAssets.ReadDir(".")
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
		rawImg, err := pngAssets.Open(f.Name())
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
