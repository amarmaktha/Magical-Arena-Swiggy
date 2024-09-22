package service

import "fmt"

// Fight initiates the fight between the two players
func (arena *Arena) Fight() {
	fmt.Println("The battle begins!")

	for arena.PlayerA.IsAlive() && arena.PlayerB.IsAlive() {
		arena.PlayTurn()
	}

	arena.DeclareWinner()
}