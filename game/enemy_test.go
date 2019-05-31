package game

import "testing"

func TestEnemyCreation(t *testing.T) {
	opponentService := opponentService{}
	opponentService.createRandomAtributes(15)

}
