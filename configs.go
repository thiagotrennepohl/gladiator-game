package main

import "github.com/thiagotrennepohl/gladiator-game/models"

var (
	warriorStatsRules = models.WarriorRules{
		Strength: 0.5,
		Stamina:  0.4,
		Dextery:  0.4,
		HP:       10,
	}

	warriorAttributesRules = models.BaseAttributes{
		Strength:     0.7,
		Stamina:      0.25,
		Dextery:      0.05,
		Intelligence: 0,
	}
)
