package beast

import (
	"fmt"
	"strconv"
)

type Beast struct {
	IsStarter    bool `yaml:"starter"`
	MainType     string
	SubType      string
	BaseStats    Stats    `yaml:"baseStats"`
	StarterMoves []string `yaml:"starterMoves"`
}

type Stats struct {
	HP    int `yaml:"hp"`
	Att   int
	Def   int
	SpAtt int `yaml:"spAtt"`
	SpDef int `yaml:"spDef"`
	Spd   int
}

type Pet struct {
	Name    string
	base    Stats
	current Stats
	Moves   []string
}

func (p Pet) Att() int {
	return p.base.Att * (p.current.Att + 2) / 2
}

func (p Pet) Def() int {
	return p.base.Def * (p.current.Def + 2) / 2
}

func (p Pet) Hp() int {
	return p.current.HP
}

func (p Pet) ModDef(v int) Pet {
	p.current.Def += v
	if p.current.Def > 6 {
		p.current.Def = 6
		return p
	}
	if p.current.Def > -6 {
		return p
	}
	p.current.Def = -6
	return p
}

func (p Pet) ModAtt(v int) Pet {
	p.current.Att += v
	if p.current.Att > 6 {
		p.current.Att = 6
		return p
	}
	if p.current.Att > -6 {
		return p
	}
	p.current.Att = -6
	return p
}

func (p Pet) OffsetHP(v int) Pet {
	p.current.HP += v
	if p.current.HP > p.base.HP {
		p.current.HP = p.base.HP
		return p
	}
	if p.current.HP >= 0 {
		return p
	}
	p.current.HP = 0
	return p
}

func (p Pet) String() string {
	m := 31
	r := p.current.HP * m / p.base.HP
	b := makeBar(r, m)
	b = fillRatio(b, p.current.HP, p.base.HP)
	return fmt.Sprintf("%s (%s)",
		p.Name, string(b))
}

func makeBar(h, m int) []rune {
	b := make([]rune, m)
	for i := 0; i < m-h; i++ {
		b[i] = ' '
	}
	for i := m - h; i < m; i++ {
		b[i] = '='
	}
	return b
}

func fillRatio(b []rune, f, s int) []rune {
	halfWay := len(b)/2 + 1
	b[halfWay] = '/'
	for j, r := range []rune(strconv.Itoa(s)) {
		b[halfWay+1+j] = r
	}
	rf := []rune(strconv.Itoa(f))
	lf := len(rf)
	for j, r := range rf {
		b[halfWay-lf+j] = r
	}
	return b
}
