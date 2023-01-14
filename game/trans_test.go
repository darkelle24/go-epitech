package game_test

import (
	"testing"

	"github.com/darkelle24/go-epitech/game"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrans(t *testing.T) {
	var gameEnv game.Game
	assert.NotNil(t, gameEnv.CreateTranspallete("name", 0, 0), "requires map to create a trans")
	require.Nil(t, gameEnv.CreateMap(10, 10), "requires map to test trans")

	const (
		name = "trans_1"
		x    = 0
		y    = 0
	)
	transSuccess := assert.Nil(t, gameEnv.CreateTranspallete(name, x, y))
	assert.Nil(t, gameEnv.CreateColis("name", 0, 1, 200))
	assert.Nil(t, gameEnv.CreateCamion("name", 1, 1, 800, 2))
	assert.NotNil(t, gameEnv.CreateTranspallete("name", 11, 10), "out of bonds")
	assert.NotNil(t, gameEnv.CreateTranspallete("name", x, y), "cell already occupied")

	if transSuccess {
		trans := gameEnv.Transps[0]
		assert.Equal(t, name, trans.GetName())
		resx, resy := trans.GetPosition()
		assert.Equal(t, resx, x)
		assert.Equal(t, resy, y)
		assert.Equal(t, game.TRANSPALET, trans.GetType())
		assert.Equal(t, "", trans.GetStatus(), "no status yet")
		assert.NotNil(t, trans.Move(0, 1, &gameEnv.Map), "cell occupied")
		assert.Equal(t, false, trans.HasColis(), "trans has no package")
		assert.Nil(t, trans.Take(0, 1, &gameEnv.Map), "should take the package")
		assert.Equal(t, "TAKE", trans.GetStatus(), "status TAKE because it took a package")
		assert.Equal(t, true, trans.HasColis(), "trans has a package")
		assert.Nil(t, trans.Move(0, 1, &gameEnv.Map))
		assert.Equal(t, "GO", trans.GetStatus(), "status GO because it moved")
		resx, resy = trans.GetPosition()
		assert.Equal(t, resx, 0)
		assert.Equal(t, resy, 1)
		assert.Nil(t, trans.NextTurn())
		trans.Wait()
		assert.Equal(t, "WAIT", trans.GetStatus(), "status WAIT because it waited")
		assert.Nil(t, trans.NextTurn())
		assert.NotNil(t, trans.Drop(0, 0, &gameEnv.Map), "should fail because there is no truck")
		assert.Nil(t, trans.Drop(1, 1, &gameEnv.Map), "should not fail because there is a truck")
		assert.Equal(t, "LEAVE", trans.GetStatus(), "status LEAVE because it droped a package")
		assert.Nil(t, trans.NextTurn())
	}
}
