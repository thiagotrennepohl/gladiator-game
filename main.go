package main

import (
	"fmt"

	"github.com/thiagotrennepohl/gladiator-game/models"
	"github.com/thiagotrennepohl/gladiator-game/service/opponent"
	"github.com/thiagotrennepohl/gladiator-game/service/warrior"
)

func main() {

	warriorRules := models.WarriorRules{
		Strength:     0.70,
		Stamina:      0.12,
		Intelligence: 0.05,
		Dextery:      0.08,
		Luck:         0.05,
	}
	WarriorBaseAttributes := models.BaseAttributes{
		Strength:     7,
		Stamina:      7,
		Dextery:      7,
		Intelligence: 7,
		Luck:         5,
	}

	warrior := warrior.NewWarriorService(models.BaseAttributes(WarriorBaseAttributes))

	opponentService := opponent.NewOpponentService(map[string]models.Character{"warrior": &warrior},
		map[string]models.BaseAttributes{"warrior": models.BaseAttributes(warriorRules)})

	enemy, err := opponentService.CreateNewOpponent(4)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", enemy.GetBaseAttributes())
	fmt.Printf("%+v\n", WarriorBaseAttributes)
	enemyAttr := enemy.GetBaseAttributes()
	fmt.Println(enemyAttr.Dextery + enemyAttr.Intelligence + enemyAttr.Luck + enemyAttr.Stamina + enemyAttr.Strength)
	fmt.Println(WarriorBaseAttributes.Dextery + WarriorBaseAttributes.Intelligence + WarriorBaseAttributes.Luck + WarriorBaseAttributes.Stamina + WarriorBaseAttributes.Strength)
}
