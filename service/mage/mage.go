package mage

import "github.com/thiagotrennepohl/gladiator-game/models"

type mageService struct {
	name           string
	baseAttributes models.BaseAttributes
}

func NewMageService(baseAttributes models.BaseAttributes) models.Character {
	return &mageService{
		baseAttributes: baseAttributes,
	}
}

func (mageSvc *mageService) GetAttackPoints() float32 {
	return float32(mageSvc.baseAttributes.Strength)*0.4 + float32(mageSvc.baseAttributes.Strength)
}

func (mageSvc *mageService) GetDefensePoints() float32 {
	return float32(mageSvc.baseAttributes.Stamina) * 0.2 * mageSvc.baseAttributes.Stamina
}

func (mageSvc *mageService) GetName() string {
	return mageSvc.name
}

//Disabled for now, because we need also add accuracy rate
//func (mageSvc *mageService) GetEvasionRate() float32 {
//Create rule for evasion rate based on luck and dextery
//	return float32(mageSvc.baseAttributes.Dextery) * 0.
//}

func (mageSvc *mageService) GetCriticalChance() float32 {
	//Create rule for criticalChance rate based on luck and dextery
	return float32(mageSvc.baseAttributes.Dextery)
}

func (mageSvc *mageService) SetName(name string) {
	//Create rule for criticalChance rate based on luck and dextery
	mageSvc.name = name
}

func (mageSvc *mageService) SetBaseAttributes(attr models.BaseAttributes) {
	mageSvc.baseAttributes = attr
}

func (mageSvc *mageService) GetBaseAttributes() models.BaseAttributes {
	return mageSvc.baseAttributes
}
