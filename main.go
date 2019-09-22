package main

import (
	"fmt"

	"github.com/thiagotrennepohl/gladiator-game/models"
	"github.com/thiagotrennepohl/gladiator-game/service/warrior"
)

func main() {

	warrior := warrior.NewWarriorService(models.BaseAttributes{
		Dextery:      3,
		Intelligence: 2,
		Luck:         0,
		Stamina:      6,
		Strength:     10,
	})
	fmt.Println(warrior.GetAttackPoints())

	// warrior := &models.Warrior{}
	// opponent := models.Character{}
	// character := models.Character{
	// 	FightStyle: warrior,
	// }
	// character.FightStyle.CalculateAttackPoints(&character.BaseAtributes, &opponent.BaseAtributes)

	//Player selects attack
	// Get attacker attack points //character.getAttackPoints
	// Get `dfenser` Defense poits //character.GerDefensePoints
	// Get `defenser` B Evasion rate
	// Calculate Damage
	// Apply Damage to defenser
	//
}
