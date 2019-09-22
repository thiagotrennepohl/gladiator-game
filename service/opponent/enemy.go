package opponent

import (
	"github.com/thiagotrennepohl/gladiator-game/models"
)

type opponentService struct {
	name           string
	baseAttributes models.BaseAttributes
}

func NewOpponentService() opponentService {
	return opponentService{}
}

func (o *opponentService) CreateNewOpponent(fightStyles map[string]models.RandomCharacterService, currentPlayerLevel int) models.Character {
	return fightStyles["warrior"]
}

func (o *opponentService) pickFightStyle() {

}

func (o *opponentService) setAttributes() {

}

//I must have rules to define which fightStyle I#m picking
// What about creating a rules service?

//func (o *opponentService) CreateNewOpponent(playerLevel int) opponentService {
//	fightStyle := o.fightStyles[o.getRandomInt(len(o.fightStyles))]
//	baseAttributes := o.opponentService.createRandomAtributes(15)createRandomAtributes(playerLevel)
//	opponent := fightStyle
//	return fightStyle
//}
//
//func (o *opponentService) getRandomInt(maxSize int) int {
//	source := rand.NewSource(time.Now().UnixNano())
//	randomizer := rand.New(source)
//	return randomizer.Intn(maxSize + 1)
//}
//
//func (o *opponentService) createRandomAtributes(attrCount int) models.BaseAtributes {
//	var baseAttributes models.BaseAtributes
//	ptr := reflect.ValueOf(&baseAttributes)
//	value := reflect.Indirect(ptr)
//	for i := 0; i < value.NumField(); i++ {
//		attr := e.getRandomInt(attrCount)
//		value.Field(i).SetInt(int64(attr))
//		attrCount -= attr
//	}
//	return baseAttributes
//}
//
//func (o *opponentService) GetAttackPoints() float32 {
//	return float32(o.baseAttributes.Strength)
//}
//
//func (o *opponentService) GetDefensePoints() float32 {
//	return float32(o.baseAttributes.Stamina * 2)
//}
//
//func (o *opponentService) GetName() string {
//	return o.name
//}
//
//func (o *opponentService) GetEvasionRate() float32 {
//	//Create rule for evasion rate based on luck and dextery
//	return float32(o.baseAttributes.Dextery)
//}
//
//func (o *opponentService) GetCriticalChance() float32 {
//	//Create rule for criticalChance rate based on luck and dextery
//	return float32(o.baseAttributes.Dextery)
//}
