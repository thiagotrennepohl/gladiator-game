package models

type BattleService interface {
}

type Character interface {
	GetName() string
	GetAttackPoints() float32
	GetDefensePoints() float32
	GetEvasionRate() float32
	GetCriticalChance() float32
	SetName(string)
	SetBaseAttributes(BaseAttributes)
	GetBaseAttributes() BaseAttributes
}

type OpponentService interface {
	CreateNewOpponent(classes map[string]Character, rules map[string]BaseAttributes) Character
}
