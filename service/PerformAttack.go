package service

import (
	"fmt"
	"magical-arena/store"
)

// PerformAttack simulates an attack from one player to another
func (arena *Arena) PerformAttack(attacker, defender *store.Player) {
	attackRoll := RollDice()
	defenseRoll := RollDice()

	attackDamage := attackRoll * attacker.Attack
	defense := defenseRoll * defender.Strength

	damage := attackDamage - defense
	if damage < 0 {
		damage = 0
	}

	defender.UpdateHealth(damage)

	fmt.Printf("%s attacks! (Die roll: %d) Attack Damage: %d\n", attacker.Name, attackRoll, attackDamage)
	fmt.Printf("%s defends! (Die roll: %d) Defense Strength: %d\n", defender.Name, defenseRoll, defense)
	fmt.Printf("%s takes %d damage, health is now %d\n\n", defender.Name, damage, defender.Health)
}