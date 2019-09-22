package opponent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagotrennepohl/gladiator-game/mocks"
	"github.com/thiagotrennepohl/gladiator-game/models"
	"github.com/thiagotrennepohl/gladiator-game/service/opponent"
)

func TestEnemyCreation(t *testing.T) {
	playerMock := mocks.OpponentMock{}
	chars := []models.Character{&playerMock}
	opponentService := opponent.NewOpponentService()
	opponent := opponentService.CreateNewOpponent(chars, 15)
	assert.Implements(t, (*models.Character)(nil), opponent)
}

func TestEnemyHasSameLevel(t *testing.T) {

}
