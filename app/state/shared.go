package state

import (
	"github.com/coopstools/minibeast/app/asset"
	"github.com/hajimehoshi/ebiten/v2"
)

type Shared struct {
	selCharacter        string
	availableCharacters []string

	BackGroundImages map[string]*ebiten.Image
	characterImages  map[string]map[string]*ebiten.Image
}

func (s *Shared) PullAvailable() []string {
	return s.availableCharacters
}

func (s *Shared) SetSelected(sel int) {
	s.selCharacter = s.availableCharacters[sel]
}

func (s *Shared) PullSelectedCharacter(imgName string) *ebiten.Image {
	return s.characterImages[s.selCharacter][imgName]
}

func NewShared() Shared {
	duckImages, otherImages := asset.LoadMulti(), asset.LoadAssets()

	availableCharacters := make([]string, 0, len(duckImages))
	for k, _ := range duckImages {
		availableCharacters = append(availableCharacters, k)
	}

	return Shared{
		selCharacter:        availableCharacters[0],
		availableCharacters: availableCharacters,
		BackGroundImages:    otherImages,
		characterImages:     duckImages,
	}
}
