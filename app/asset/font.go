package asset

import (
	_ "embed"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"image/draw"
)

//go:embed font.ttf
var fontFile []byte

type Builder struct {
	img        *ebiten.Image
	drawer     *freetype.Context
	lineNumber int
}

func (b *Builder) Println(value string) {
	pt := freetype.Pt(4, 12+16*b.lineNumber)
	_, _ = b.drawer.DrawString(value, pt)
	b.lineNumber += 1
}

func (b *Builder) Finish(drawImage func(img *ebiten.Image, options *ebiten.DrawImageOptions)) {
	opt := &ebiten.DrawImageOptions{}
	drawImage(b.img, opt)
}

func LoadFont() *truetype.Font {
	loadedFont, _ := truetype.Parse(fontFile)
	return loadedFont
}

func NewBuilder(numLines int) Builder {
	width, height := 300, 16*numLines

	loadedFont, err := truetype.Parse(fontFile)
	if err != nil {
		panic(err)
	}

	opts := truetype.Options{}
	opts.Size = 8
	//face := truetype.NewFace(loadedFont, &opts)

	// sampling of some of the options that are set
	a := uint8(50)
	fg, bg := image.Black, image.NewUniform(color.RGBA{R: a, G: a, B: a, A: a})
	rgba := ebiten.NewImage(width, height)
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetFont(loadedFont)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	return Builder{
		img:        rgba,
		drawer:     c,
		lineNumber: 0,
	}
}
