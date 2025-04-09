package models

type TileType string

const (
	TileGrass0 TileType = "grass_tile_0"
	TileGrass1 TileType = "grass_tile_1"
	TileFlower TileType = "flowers_tile"
	// Add more tile types as needed
)

type TileInfo struct {
	OriginalType TileType
	Type         TileType
	Rotation     uint8
}
