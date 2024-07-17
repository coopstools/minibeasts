package display

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"image/draw"
)

type popup1col struct {
	font  *truetype.Font
	lines []string
}

func (p popup1col) AddLine(value string) popup1col {
	p.lines = append(p.lines, value)
	return p
}

func (p popup1col) Draw() *ebiten.Image {
	width, height := 300, 16*(len(p.lines)+1)

	opts := truetype.Options{}
	opts.Size = 8

	a := uint8(50)
	fg, bg := image.Black, image.NewUniform(color.RGBA{R: a, G: a, B: a, A: a})
	rgba := ebiten.NewImage(width, height)
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetFont(p.font)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	for i, line := range p.lines {
		pt := freetype.Pt(4, 12+16*i)
		_, _ = c.DrawString(line, pt)
	}

	return rgba
}
