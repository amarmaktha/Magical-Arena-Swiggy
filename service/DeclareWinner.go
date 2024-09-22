package service

import "fmt"

func (arena *Arena) DeclareWinner() {
	if !arena.PlayerA.IsAlive() {
		fmt.Printf("%s has been defeated! %s wins the battle.\n", arena.PlayerA.Name, arena.PlayerB.Name)
	} else {
		fmt.Printf("%s has been defeated! %s wins the battle.\n", arena.PlayerB.Name, arena.PlayerA.Name)
	}
}