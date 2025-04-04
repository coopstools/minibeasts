package tiles

import (
	"bytes"
	"embed"
	"image"
	"image/png"
	"sync"

	"github.com/pkg/errors"
)

//go:embed *.png
var tileset embed.FS

var Load = sync.OnceValues(load)

func load() (map[string]image.Image, error) {
	tiles := make(map[string]image.Image)

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
		tiles[file.Name()[:len(file.Name())-4]] = img
	}

	return tiles, nil
}
