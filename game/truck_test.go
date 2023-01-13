package game_test

import (
	"testing"

	"github.com/darkelle24/go-epitech/game"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTruck(t *testing.T) {
	var gameEnv game.Game
	assert.NotNil(t, gameEnv.CreateCamion("name", 0, 0, 1000, 5), "requires map to create a truck")
	require.Nil(t, gameEnv.CreateMap(10, 10), "requires map to test truck")

	const (
		name      = "truck_1"
		maxWeight = 1000
		x         = 0
		y         = 0
	)
	truckSuccess := assert.Nil(t, gameEnv.CreateCamion(name, x, y, maxWeight, 2))
	assert.NotNil(t, gameEnv.CreateCamion("name", 11, 10, 1000, 5), "out of bonds")
	assert.NotNil(t, gameEnv.CreateCamion("name", x, y, 1000, 5), "cell already occupied")
	assert.NotNil(t, gameEnv.CreateCamion("name", x+1, y, -1000, 5), "negative weight")
	assert.NotNil(t, gameEnv.CreateCamion("name", x+2, y, 1000, -5), "negative turns")

	if truckSuccess {
		truck := gameEnv.Trucks[0]
		assert.Equal(t, 2, truck.GetTimeMax())
		assert.Equal(t, name, truck.GetName())
		assert.Equal(t, 0, truck.GetCurrentWeight())
		assert.Equal(t, maxWeight, truck.GetMaxWeight())
		resx, resy := truck.GetPosition()
		assert.Equal(t, resx, x)
		assert.Equal(t, resy, y)
		assert.Equal(t, game.CAMION, truck.GetType())
		assert.Equal(t, "", truck.GetStatus(), "no status yet")
		assert.Equal(t, true, truck.IsPresent(), "truck hasn't moved yet")
		pack := game.Colis{}
		assert.Nil(t, truck.AddPackage(&pack), "should add the package")
		// can't check if package is added because Colis weights 0
		assert.Nil(t, truck.Move(0, 0, &gameEnv.Map))
		assert.Equal(t, "GONE", truck.GetStatus(), "status GONE because it moved")
		assert.Equal(t, false, truck.IsPresent(), "truck has moved")
		assert.Nil(t, truck.NextTurn())
		assert.Equal(t, false, truck.IsPresent(), "truck has yet to return 1")
		assert.Equal(t, 2, truck.GetComeback(), "comeback should equal 2")
		assert.Nil(t, truck.NextTurn())
		assert.Equal(t, false, truck.IsPresent(), "truck has yet to return 2")
		assert.Equal(t, 1, truck.GetComeback(), "comeback should equal 1")
		assert.NotNil(t, truck.Move(0, 0, &gameEnv.Map), "truck is GONE, shoudn't be able to move")
		assert.Nil(t, truck.NextTurn())
		assert.Equal(t, "", truck.GetStatus(), "no status because new turn")
		assert.Equal(t, true, truck.IsPresent(), "truck should be back")
		assert.Nil(t, truck.Wait())
		assert.NotNil(t, truck.Wait(), "can do only one action per turn")
		assert.Nil(t, truck.NextTurn())
		if assert.Equal(t, "", truck.GetStatus()) {
			assert.NotNil(t, truck.NextTurn(), "should fail because no actions were taken this turn")
		}
	}
}
