package asset

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"io/fs"
)

//go:embed png/*.png
var pngAssets embed.FS

//go:embed png/multi/*.png
var multiAssets embed.FS

func LoadAssets() map[string]*ebiten.Image {
	fmt.Println("Loading Assets")
	dir, err := pngAssets.ReadDir("png")
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
		rawImg, err := pngAssets.Open("png/" + f.Name())
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

func LoadMulti() map[string]map[string]*ebiten.Image {
	dirName := "png/multi"
	dir, err := multiAssets.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	allImages := make(map[string]map[string]*ebiten.Image, len(dir))
	for _, asset := range dir {
		nl := len(asset.Name())
		if nl <= 4 {
			continue
		}
		if asset.Name()[nl-4:] != ".png" {
			continue
		}
		assetName := asset.Name()[:nl-4]
		rawImg, err := multiAssets.Open(dirName + "/" + asset.Name())
		if err != nil {
			panic(err)
		}
		allImages[assetName] = splitImage(rawImg)
	}
	return allImages
}

func splitImage(imgFs fs.File) map[string]*ebiten.Image {
	img, _, err := ebitenutil.NewImageFromReader(imgFs)
	if err != nil {
		panic(err)
	}
	imgNames := []string{"side1", "side2", "down", "up", "profile", ""}
	bounds := img.Bounds().Max
	cols, rows := (bounds.X+1)/17, (bounds.Y+1)/17
	images := make(map[string]*ebiten.Image, cols*rows+1)
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			x0, y0, x1, y1 := 17*i, 17*j, 17*i+16, 17*j+16
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1))
			imgName := imgNames[j+i*rows]
			images[imgName] = ebiten.NewImageFromImage(subImg)
		}
	}
	return images
}
