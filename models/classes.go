package models

//Character represents a player
type Character struct {
	Name          string
	Gender        string
	BaseAtributes BaseAtributes
	FightStyle    FightStyle
	//Inventory     []Item
}

//FightStyle represents the base contract of a new class
type FightStyle interface {
	CalculateAttackPoints(playerAtributes, opponentAtribute *BaseAtributes) float32
	CalculateDefensePoints(atributes *BaseAtributes) float32
	HasEvaded(atributes *BaseAtributes) bool
	CalculateCriticalChance(atributes *BaseAtributes) float32
	HasFled(atributes *BaseAtributes) bool
}

//BaseAtributes represents basic character atributes set up
type BaseAtributes struct {
	Strength     int
	Stamina      int
	Dextery      int
	Intelligence int
	Luck         int
}

type Warrior struct {
}

type Archer struct {
}

func (w *Warrior) CalculateAttackPoints(playerAtribute, opponentAtribute *BaseAtributes) float32 {
	return float32(playerAtribute.Strength/10 - opponentAtribute.Stamina/8)
}
func (w *Warrior) CalculateDefensePoints(atribute *BaseAtributes) float32 {
	return float32(atribute.Stamina / 8)
}

func (w *Warrior) HasEvaded(atributes *BaseAtributes) bool {
	//não faço ideia
	return true
}

func (w *Warrior) CalculateCriticalChance(atributes *BaseAtributes) float32 {
	//não faço ideia
	return 1
}

func (w *Warrior) HasFled(atributes *BaseAtributes) bool {

func (a *Archer) CalculateAttackPoints(playerAtribute, opponentAtribute *BaseAtributes) float32 {
	return float32(playerAtribute.Dextery/10 - opponentAtribute.Stamina/8)
}

func (a *Archer) CalculateDefensePoints(atribute *BaseAtributes) float32 {
	return float32(atribute.Stamina / 8)
}
func (a *Archer) HasEvaded(atribute *BaseAtributes) bool {
	return true //float32(atribute.Dextery / 7)
}

func (a *Archer) CalculateCriticalChance(atribute *BaseAtributes) float32 {
	return float32(atribute.Luck / 6)
}

func (a *Archer) HasFled(atribute *BaseAtributes) bool {

	return true
}

type Item interface {
}

//Effects should always be processed before the player's turn.

//Use -> returns an effect.
//How will I know what this effect does?
//Or how will I apply this effect to the atributes? Dparam?
//A giant switch case where depending on the type of the effect I'll apply to the user atributes
//But where should this logic be?
//Type of effects
//if player.HasEffect -> case effect type of Poison -> damage life overtime (duration) ->
//The player has to have a Effect atribute
//will be the battle responsible for controlling this?
//Quando o jogador usa um evento da vida, dispara um channel que faz esse
//swith case maldito e aplica os eventos
//Channel para fazer isso - QUal a melhor forma de fazer?
//o personagem precisa der "metodos" para interagir com esse canal.
type Consumable struct {
}

type Wearable struct {
}
