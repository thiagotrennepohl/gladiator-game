package warrior

import (
	"github.com/thiagotrennepohl/gladiator-game/models"
)

type warriorService struct {
	name           string
	baseAttributes models.BaseAttributes
	statsModifier  models.BaseAttributes
}

//NewWarriorService creates a new warrior
func NewWarriorService(attributesRules models.BaseAttributes) models.Character {
	return &warriorService{
		statsModifier: attributesRules,
	}
}
func (warriorSvc *warriorService) GetAttackPoints() float32 {
	return warriorSvc.baseAttributes.Strength*warriorSvc.statsModifier.Strength + 10
}

func (warriorSvc *warriorService) GetDefensePoints() float32 {
	return warriorSvc.baseAttributes.Stamina*warriorSvc.statsModifier.Stamina + 10
}

func (warriorSvc *warriorService) GetName() string {
	return warriorSvc.name
}

//Disabled for now, because we need also add accuracy rate
//func (warriorSvc *warriorService) GetEvasionRate() float32 {
//Create rule for evasion rate based on luck and dextery
//	return float32(warriorSvc.baseAttributes.Dextery) * 0.
//}

func (warriorSvc *warriorService) GetCriticalChance() float32 {
	//Create rule for criticalChance rate based on luck and dextery
	return float32(warriorSvc.baseAttributes.Dextery)
}

func (warriorSvc *warriorService) SetName(name string) {
	//Create rule for criticalChance rate based on luck and dextery
	warriorSvc.name = name
}

func (warriorSvc *warriorService) SetBaseAttributes(attr models.BaseAttributes) {
	warriorSvc.baseAttributes = attr
}

func (warriorSvc *warriorService) GetBaseAttributes() models.BaseAttributes {
	return warriorSvc.baseAttributes
}

func (warriorSvc *warriorService) GetHealthPoints() float32 {
	return warriorSvc.baseAttributes.HP
}

func (warriorSvc *warriorService) DecreaseFromHealthPoints(dmg float32) {
	warriorSvc.baseAttributes.HP -= dmg
}
