package game

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/thiagotrennepohl/gradiator-game/models"
)

type opponentService struct {
	fightStyles []models.FightStyle
}

//NewOpponentService creates a new instance of opponentService for creating opponents
//as the player gets stronger
func NewOpponentService(fightStyles []models.FightStyle) opponentService {
	return opponentService{
		fightStyles: fightStyles,
	}
}

func (e *opponentService) CreateOpponent(playerLevel int) models.Character {
	fightStyle := e.fightStyles[e.getRandomInt(len(e.fightStyles))]
	baseAttributes := e.createRandomAtributes(playerLevel)
	opponent := models.Character{
		FightStyle:    fightStyle,
		BaseAtributes: baseAttributes,
	}
	return opponent
}

func (e *opponentService) getRandomInt(maxSize int) int {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	return randomizer.Intn(maxSize + 1)
}

func (e *opponentService) createRandomAtributes(attrCount int) models.BaseAtributes {
	var baseAttributes models.BaseAtributes
	ptr := reflect.ValueOf(&baseAttributes)
	value := reflect.Indirect(ptr)
	for i := 0; i < value.NumField(); i++ {
		attr := e.getRandomInt(attrCount)
		value.Field(i).SetInt(int64(attr))
		attrCount -= attr
	}
	return baseAttributes
}
