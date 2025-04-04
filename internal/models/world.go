package models

import (
	"math/rand"
)

type WorldData struct {
	Tiles [][]TileInfo
	// We'll add sprites and other persistent data here later
}

func NewWorldData(width, height int) *WorldData {
	tiles := make([][]TileInfo, height)
	for y := range tiles {
		tiles[y] = make([]TileInfo, width)
		for x := range tiles[y] {
			// The tiles we will be using are grass_tile_0, grass_tile_1, and flower_tile
			// Randomly assign tile type and rotation

			tileType := TileGrass0
			r := rand.Float64()
			if r < .6 { // 40% chance of grass_tile_1 (80% change for any grass tile)
				tileType = TileGrass1
			}
			if r < 0.2 { // 20% chance of flowers
				tileType = TileFlower
			}
			rotation := rand.Intn(4) * 90 // 0, 90, 180, or 270 degrees

			tiles[y][x] = TileInfo{
				Type:     tileType,
				Rotation: rotation,
			}
		}
	}

	return &WorldData{
		Tiles: tiles,
	}
}
