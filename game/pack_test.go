package game_test

import (
	"testing"

	"github.com/darkelle24/go-epitech/game"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPackage(t *testing.T) {
	var gameEnv game.Game
	assert.NotNil(t, gameEnv.CreateColis("name", 0, 0, 200), "requires map to create a package")
	require.Nil(t, gameEnv.CreateMap(10, 10), "requires map to test package")

	const (
		name   = "pack_1"
		x      = 0
		y      = 0
		weight = 200
	)
	colisSuccess := assert.Nil(t, gameEnv.CreateColis(name, x, y, weight))
	assert.NotNil(t, gameEnv.CreateColis("name", 11, 10, 200), "out of bonds")
	assert.NotNil(t, gameEnv.CreateColis("name", x, y, 200), "cell already occupied")
	assert.NotNil(t, gameEnv.CreateColis("name", x+1, y+1, -100), "negative weight")

	if colisSuccess {
		tile := gameEnv.Map[x][y]
		colis, ok := tile.Tool.(*game.Colis)
		assert.Equal(t, true, ok)
		if ok {
			assert.Equal(t, name, colis.GetName())
			resx, resy := colis.GetPosition()
			assert.Equal(t, resx, x)
			assert.Equal(t, resy, y)
			assert.Equal(t, game.COLIS, colis.GetType())
			assert.Equal(t, false, colis.IsDelivered(), "not delivered yet")
			colis.SetDelivered()
			assert.Equal(t, true, colis.IsDelivered(), "is delivered")
			assert.Equal(t, weight, colis.GetCurrentWeight())
			assert.Equal(t, weight, colis.GetMaxWeight())
		}
	}
}
