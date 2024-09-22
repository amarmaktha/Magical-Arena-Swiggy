package service

// this func will decide which player has to attack  and change the turn of player
func (arena *Arena) PlayTurn() {
	if arena.PlayerA.Health <= arena.PlayerB.Health {
		arena.PerformAttack(arena.PlayerA, arena.PlayerB)
	} else {
		arena.PerformAttack(arena.PlayerB, arena.PlayerA)
	}
}
