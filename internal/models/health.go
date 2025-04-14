package models

import "time"

type Health struct {
	Current      int
	Max          int
	LastHitTime  time.Time
	ImmunityTime time.Duration
}

func NewHealth(vitality int) *Health {
	// Base health of 20, each point of vitality adds 5 health
	maxHealth := 20 + (vitality-8)*5
	return &Health{
		Current:      maxHealth,
		Max:          maxHealth,
		ImmunityTime: time.Second, // 1 second immunity after being hit
	}
}

func (h *Health) CanTakeDamage() bool {
	return time.Since(h.LastHitTime) >= h.ImmunityTime
}

func (h *Health) TakeDamage(amount int) {
	if !h.CanTakeDamage() {
		return
	}
	h.Current -= amount
	if h.Current < 0 {
		h.Current = 0
	}
	h.LastHitTime = time.Now()
}

func (h *Health) Heal(amount int) {
	h.Current += amount
	if h.Current > h.Max {
		h.Current = h.Max
	}
}
