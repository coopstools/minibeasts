package models

type TileType string

const (
	TileGrass0 TileType = "grass_tile_0"
	TileGrass1 TileType = "grass_tile_1"
	TileFlower TileType = "flower_tile"
	// Add more tile types as needed
)

type TileInfo struct {
	Type     TileType
	Rotation int
}
