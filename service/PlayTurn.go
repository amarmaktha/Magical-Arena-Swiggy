package service

// this func will play the turn of the players
func (arena *Arena) PlayTurn() {
	if arena.PlayerA.Health <= arena.PlayerB.Health {
		arena.PerformAttack(arena.PlayerA, arena.PlayerB)
	} else {
		arena.PerformAttack(arena.PlayerB, arena.PlayerA)
	}
}