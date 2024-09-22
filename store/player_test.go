package store

import "testing"

func TestUpdateHealth(t *testing.T) {
	player := &Player{Name: "Test Player", Health: 100, Strength: 5, Attack: 10}
	player.UpdateHealth(30)

	if player.Health != 70 {
		t.Errorf("Expected health to be 70, got %d", player.Health)
	}

	player.UpdateHealth(80)
	if player.Health != 0 {
		t.Errorf("Expected health to be 0, got %d", player.Health)
	}
}

func TestIsAlive(t *testing.T) {
	player := &Player{Name: "Test Player", Health: 50}
	if !player.IsAlive() {
		t.Errorf("Player should be alive")
	}

	player.Health = 0
	if player.IsAlive() {
		t.Errorf("Player should be dead")
	}
}
