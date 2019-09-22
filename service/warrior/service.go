package warrior

import "github.com/thiagotrennepohl/gladiator-game/models"

type warriorService struct {
	name           string
	baseAttributes models.BaseAttributes
}

//NewWarriorService creates a new warrior
func NewWarriorService(attributes models.BaseAttributes) warriorService {
	return warriorService{
		baseAttributes: attributes,
	}
}
func (warriorSvc *warriorService) GetAttackPoints() float32 {
	return float32(warriorSvc.baseAttributes.Strength)
}

func (warriorSvc *warriorService) GetDefensePoints() float32 {
	return float32(warriorSvc.baseAttributes.Stamina * 2)
}

func (warriorSvc *warriorService) GetName() string {
	return warriorSvc.name
}

func (warriorSvc *warriorService) GetEvasionRate() float32 {
	//Create rule for evasion rate based on luck and dextery
	return float32(warriorSvc.baseAttributes.Dextery)
}

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
