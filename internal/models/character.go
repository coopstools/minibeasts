package models

type Character struct {
	Name       string
	Profession string
	Strength   int
	Dexterity  int
	Vitality   int
	// Add more stats as needed
}

func NewCharacter() *Character {
	return &Character{
		Strength:  10,
		Dexterity: 10,
		Vitality:  10,
	}
}
