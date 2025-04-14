package models

import (
	"math/rand"
)

type Character struct {
	Name         string
	Profession   string
	Strength     int
	Dexterity    int
	Vitality     int
	Intelligence int
	StatPool     int
	Health       *Health
}

const (
	MinStatValue    = 8
	BaseStatValue   = 10
	InitialStatPool = 2
)

func NewCharacter() *Character {
	c := &Character{
		Name:         generateRandomName(),
		Profession:   "Warrior",
		Strength:     BaseStatValue,
		Dexterity:    BaseStatValue,
		Vitality:     BaseStatValue,
		Intelligence: BaseStatValue,
		StatPool:     InitialStatPool,
	}
	c.Health = NewHealth(c.Vitality)
	return c
}

func NewMinCharacter() *Character {
	c := NewCharacter()
	c.Strength = MinStatValue
	c.Dexterity = MinStatValue
	c.Vitality = MinStatValue
	c.Intelligence = MinStatValue
	c.UpdateStats()
	return c
}

func (c *Character) UpdateStats() {
	c.Health = NewHealth(c.Vitality)
}

func generateRandomName() string {
	prefixes := []string{"Brave", "Swift", "Wise", "Strong", "Clever"}
	suffixes := []string{"walker", "smith", "heart", "soul", "mind"}

	return prefixes[rand.Intn(len(prefixes))] + suffixes[rand.Intn(len(suffixes))]
}
