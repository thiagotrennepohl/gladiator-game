package opponent

import (
	"fmt"

	"github.com/thiagotrennepohl/gladiator-game/models"
)

type opponentService struct {
	classes      map[string]models.Character
	classesRules map[string]models.BaseAttributes
}

//NewOpponentService creates a new implementation of the OpponentService interface
func NewOpponentService(classes map[string]models.Character, classesRules map[string]models.BaseAttributes) models.OpponentService {
	return &opponentService{
		classes:      classes,
		classesRules: classesRules,
	}
}

//CreateNewOpponent creates a new opponent based on the players level
func (o *opponentService) CreateNewOpponent(currentPlayerLevel int) (models.Character, error) {
	className, character := o.pickFightStyle()
	baseAttributes, err := o.setAttributes(className, currentPlayerLevel)
	if err != nil {
		return nil, err
	}
	character.SetBaseAttributes(baseAttributes)
	return character, err
}

//pickFightStyle relies on Golang Map shuffling.
//https://blog.golang.org/go-maps-in-action
func (o *opponentService) pickFightStyle() (className string, fightStyle models.Character) {
	if len(o.classes) >= 1 {
		for className, fightStyle := range o.classes {
			return className, fightStyle
		}
	}
	return
}

func (o *opponentService) setAttributes(className string, playerLevel int) (models.BaseAttributes, error) {
	if rules, ok := o.classesRules[className]; ok {
		//5 base points for each attribute, plus the points for each level
		//If we increase the number of attributes we'll need to change this implementation.
		basePoints := playerLevel * 2
		return models.BaseAttributes{
			Strength:     float32(basePoints)*rules.Strength + 5,
			Stamina:      float32(basePoints)*rules.Stamina + 5,
			Dextery:      float32(basePoints)*rules.Dextery + 5,
			Intelligence: float32(basePoints)*rules.Intelligence + 5,
			Luck:         float32(basePoints)*rules.Luck + 5,
		}, nil
	}
	return models.BaseAttributes{}, fmt.Errorf("Class %s not found in the rules list", className)
}
