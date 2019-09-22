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
	SetAttackPoints(float32)
	CreateRandomCharacter() Character
}

type OpponentService interface {
	CreateNewOpponent(fightStyles []Character, playerLevel int) Character
}
