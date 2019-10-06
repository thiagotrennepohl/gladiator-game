package models

type Character interface {
	GetName() string
	GetAttackPoints() float32
	GetDefensePoints() float32
	//GetEvasionRate() float32
	GetCriticalChance() float32
	SetName(string)
	SetBaseAttributes(BaseAttributes)
	GetBaseAttributes() BaseAttributes
	GetHealthPoints() float32
	DecreaseFromHealthPoints(float32)
}

type OpponentService interface {
	CreateNewOpponent(currentPlayerLevel int) (Character, error)
}

type MatchService interface {
	IsInProgress() bool
	StartMatch() error
	SkipTurn()
	TossCoin() bool
}
type CharacterCreationService interface {
	CreateCharacter(name string, class string, attributes BaseAttributes) Character
	GetAvailableClasses() (map[string]Character, error)
}
