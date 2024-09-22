package handler

import (
	"magical-arena/service"
	"magical-arena/store"
)

// StartGame handler is used to initialize players and starts the game.
// Used dummy values for health and strength and attack.
func StartGame() {

	playerA := &store.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
	playerB := &store.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}

	arena := &service.Arena{PlayerA: playerA, PlayerB: playerB}

	arena.Fight()
}
