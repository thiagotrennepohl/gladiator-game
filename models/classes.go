package models

//BaseAtributes represents basic character atributes set up
type BaseAttributes struct {
	Strength     int
	Stamina      int
	Dextery      int
	Intelligence int
	Luck         int
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
