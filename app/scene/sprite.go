package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// A B C D  E F G H
// | | | |  | | | step
// | | | |  | | vertical
// | | | |  | up/right
const (
	dirUp    byte = 6
	dirDown  byte = 2
	dirLeft  byte = 0
	dirRight byte = 4

	speed = 2
)

type Sprite struct {
	x16 int
	y16 int

	vx16 int
	vy16 int

	dir  byte
	Imgs []string
}

func (s *Sprite) MoveRight() {
	s.dir = s.dir&0xF9 + dirRight
	switch {
	case s.vx16 <= 0:
		s.vx16 = speed
	case s.vx16 <= 8:
		s.vx16 += speed
	}
	s.x16 += s.vx16
	if s.x16 > 3200 {
		s.x16 = 3200
	}
}

func (s *Sprite) MoveLeft() {
	s.dir = s.dir&0xF9 + dirLeft
	switch {
	case s.vx16 >= 0:
		s.vx16 = -speed
	case s.vx16 >= -8:
		s.vx16 -= speed
	}
	s.x16 += s.vx16
	if s.x16 < 0 {
		s.x16 = 0
	}
}

func (s *Sprite) MoveDown() {
	s.dir = s.dir&0xF9 + dirDown
	switch {
	case s.vy16 <= 0:
		s.vy16 = speed
	case s.vy16 <= 8:
		s.vy16 += speed
	}
	s.y16 += s.vy16
	if s.y16 > 3200 {
		s.y16 = 3200
	}
}

func (s *Sprite) MoveUp() {
	s.dir = s.dir&0xF9 + dirUp
	switch {
	case s.vy16 >= 0:
		s.vy16 = -speed
	case s.vy16 >= -8:
		s.vy16 -= speed
	}
	s.y16 += s.vy16
	if s.y16 < 0 {
		s.y16 = 0
	}
}

func (s *Sprite) ResetVelocity() {
	s.vx16, s.vy16 = 0, 0
}

func (s *Sprite) UpdateLegs() {
	if s.vy16 == 0 && s.vx16 == 0 {
		return
	}
	s.dir ^= 1
}

func (s *Sprite) DisplayNameAndOpt() (string, *ebiten.DrawImageOptions) {
	scaleX, scaleY := 4.0, 4.0
	transX, transY := float64(s.x16/8.0), float64(s.y16/8.0)
	var imgName string
	switch s.dir & 7 {
	case 0b000: //000
		//horizontal and left, step1
		imgName = "side1"
		scaleX *= -1.0
		transX += 64.
	case 0b001:
		//horizontal and left, step2
		imgName = "side2"
		scaleX *= -1.0
		transX += 64.00
	case 0b010:
		//vertical and down, step1
		imgName = "down"
	case 0b011:
		//vertical and down, step2
		imgName = "down"
		scaleX *= -1.0
		transX += 64.0
	case 0b100:
		//horizontal and right, step1
		imgName = "side1"
	case 0b101:
		//horizontal and right, step2
		imgName = "side2"
	case 0b110:
		//vertical and up, step1
		imgName = "up"
	case 0b111:
		//vertical and up, step2
		imgName = "up"
		scaleX *= -1.0
		transX += 64.0
	}

	opt := ebiten.DrawImageOptions{}
	opt.GeoM.Scale(scaleX, scaleY)
	opt.GeoM.Translate(transX, transY)

	return imgName, &opt
}
