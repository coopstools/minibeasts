package battle

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

type Popup2x2 struct {
	font *truetype.Font
	opts []string
	sel  byte
	xPos int
	yPos int
}

func (p *Popup2x2) MoveLeft() {
	p.sel ^= 1
}

func (p *Popup2x2) MoveUp() {
	p.sel ^= 2
}

func (p *Popup2x2) Draw() *ebiten.Image {
	width, height := 300, 32

	opts := truetype.Options{}
	opts.Size = 8

	a := uint8(100)
	fg, bg := image.Black, image.NewUniform(color.RGBA{R: a, G: a, B: a, A: a})
	rgba := ebiten.NewImage(width, height)
	draw.Draw(rgba, rgba.Bounds(), bg, image.Pt(0, 0), draw.Src)
	c := freetype.NewContext()
	c.SetFont(p.font)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	for i, pt := range []fixed.Point26_6{
		freetype.Pt(4, 12), freetype.Pt(150, 12),
		freetype.Pt(4, 28), freetype.Pt(150, 28),
	} {
		//pt := freetype.Pt(4, 12+16*i)
		prefix := " "
		if i == int(p.sel) {
			prefix = ">"
		}
		_, _ = c.DrawString(prefix+p.opts[i], pt)
	}
	return rgba
}

func (p *Popup2x2) PullCurrent() int {
	return int(p.sel)
}

func New2x2Popup(font *truetype.Font, opts []string, x, y int) Popup2x2 {
	return Popup2x2{
		font: font,
		opts: opts,
		xPos: x,
		yPos: y,
	}
}
