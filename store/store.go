package store

// struct to represent a player.
type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}
// Actually instead of the below functions we can use Dbs for uodate health and check if the person is alive or not
// Due to time constraints i am using dummy functions.

// this function is used to update the health of the player
func (p *Player) UpdateHealth(damage int) {
	p.Health -= damage
	if p.Health < 0 {
		p.Health = 0
	}
}

// Check if the player is still alive or not
func (p *Player) IsAlive() bool {
	return p.Health > 0
}
