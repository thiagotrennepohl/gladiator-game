package main

import "github.com/thiagotrennepohl/gradiator-game/models"

func main() {
	warrior := &models.Warrior{}
	opponent := models.Character{}
	character := models.Character{
		FightStyle: warrior,
	}
	character.FightStyle.CalculateAttackPoints(&character.BaseAtributes, &opponent.BaseAtributes)
}
