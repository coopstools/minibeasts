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
	// Add more stats as needed
	StatPool int
}

const (
	MinStatValue    = 8
	InitialStatPool = 2
)

func NewCharacter() *Character {
	return &Character{
		Name:         generateRandomName(),
		Profession:   "Warrior",
		Strength:     10,
		Dexterity:    10,
		Vitality:     10,
		Intelligence: 10,
		StatPool:     InitialStatPool,
	}
}

func generateRandomName() string {
	prefixes := []string{"Brave", "Swift", "Wise", "Strong", "Clever"}
	suffixes := []string{"walker", "smith", "heart", "soul", "mind"}

	return prefixes[rand.Intn(len(prefixes))] + suffixes[rand.Intn(len(suffixes))]
}
