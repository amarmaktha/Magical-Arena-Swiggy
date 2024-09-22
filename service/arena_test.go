package service

import (
	"magical-arena/store"
	"testing"
)

func TestRollDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		roll := RollDice()
		if roll < 1 || roll > 6 {
			t.Errorf("RollDice produced an invalid roll: %d", roll)
		}
	}
}

func TestPerformAttack(t *testing.T) {
	attacker := &store.Player{Name: "Attacker", Health: 100, Attack: 20, Strength: 10}
	defender := &store.Player{Name: "Defender", Health: 100, Attack: 15, Strength: 5}

	arena := &Arena{PlayerA: attacker, PlayerB: defender}

	arena.PerformAttack(attacker, defender)

	// Check that health was reduced (as we don't know exact dice roll, just check if damage was done)
	if defender.Health >= 100 {
		t.Errorf("Expected defender health to be reduced")
	}
}
