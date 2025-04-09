package tiles

import (
	"bytes"
	"embed"
	"image/png"
	"math"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pkg/errors"
)

//go:embed *.png
var tileset embed.FS

var Load = sync.OnceValues(load)

func load() (map[string][]*ebiten.Image, error) {
	tiles := make(map[string][]*ebiten.Image)

	files, err := tileset.ReadDir(".")
	if err != nil {
		return nil, errors.Wrap(err, "failed to read tileset")
	}

	for _, file := range files {
		// All siles are .ong, and thus no need to check if directory
		imgBytes, err := tileset.ReadFile(file.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read tile %s", file.Name())
		}

		img, err := png.Decode(bytes.NewReader(imgBytes))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode tile %s", file.Name())
		}
		rotations := make([]*ebiten.Image, 4)
		for i := 0; i < 4; i++ {
			eimg := ebiten.NewImage(16, 16)
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(-8, -8)
			opt.GeoM.Rotate(float64(i) * 90 * math.Pi / 180)
			opt.GeoM.Translate(8, 8)
			eimg.DrawImage(ebiten.NewImageFromImage(img), opt)
			rotations[i] = eimg
		}
		tiles[file.Name()[:len(file.Name())-4]] = rotations
	}

	return tiles, nil
}
