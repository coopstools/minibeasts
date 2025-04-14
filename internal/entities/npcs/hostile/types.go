package hostile

type MovementState int

const (
	Idle MovementState = iota
	Wandering
	Pausing
	Pursuing
)

type NPCType int

const (
	Basic NPCType = iota
	// Add more NPC types here as needed, for example:
	// Aggressive
	// Ranged
	// Boss
)

const (
	// Movement constants
	WanderSpeed      = 1.0
	WanderRadius     = 100.0
	MinWanderTime    = 2.0
	MaxWanderTime    = 5.0
	MinPauseTime     = 1.0
	MaxPauseTime     = 3.0
	NPCAvoidRadius   = 25.0
	NPCAvoidStrength = 0.5
	MaxSpeed         = 2.0

	// Combat constants
	DefaultDetectionRange = 150.0
	DefaultDamageAmount   = 5
)
