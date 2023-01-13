package game_test

import (
	"testing"

	"github.com/darkelle24/go-epitech/game"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	var gameEnv game.Game
	assert.Nil(t, gameEnv.CreateMap(10, 10), "requires map to test package")
	gameEnv.SetTurns(2)
	assert.Equal(t, 2, gameEnv.GetTurns(), "Set to 2 turns")
	assert.Equal(t, 0, gameEnv.Turn, "current turn is 0")
	gameEnv.NextTurn()
	assert.Equal(t, 1, gameEnv.Turn, "current turn is 1")
	assert.Equal(t, "😎", gameEnv.EndStateCharacter())
	assert.Equal(t, true, gameEnv.IsAllDelivered())
	assert.Equal(t, true, gameEnv.IsDone())
	assert.Nil(t, gameEnv.CreateColis("toto", 0, 0, 200))
	assert.Equal(t, false, gameEnv.IsDone())
	assert.Equal(t, false, gameEnv.IsAllDelivered())
	assert.Equal(t, "🙂", gameEnv.EndStateCharacter())
	gameEnv.NextTurn()
	assert.Equal(t, true, gameEnv.IsDone())
}
