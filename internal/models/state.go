package models

type GameState struct {
	Character *Character
	// Add more game state here as needed
}

func NewGameState() *GameState {
	return &GameState{
		Character: NewCharacter(),
	}
}
