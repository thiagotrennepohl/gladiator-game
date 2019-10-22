package models

//BaseAtributes represents basic character atributes set up
type BaseAttributes struct {
	Strength     float32
	Stamina      float32
	Dextery      float32
	Intelligence float32
	Luck         float32
	HP           float32
	Level        int
}

type WarriorRules BaseAttributes

type MageRules BaseAttributes

type Archer BaseAttributes

type Item interface {
}

type Consumable struct {
}

type Wearable struct {
}
