package game_test

import (
	"testing"

	"github.com/darkelle24/go-epitech/game"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTruck(t *testing.T) {
	var gameEnv game.Game
	assert.NotNil(t, gameEnv.Create_camion("name", 0, 0, 1000, 5), "requires map to create a truck")
	require.Nil(t, gameEnv.Create_map(10, 10), "requires map to test truck")

	const (
		name      = "truck_1"
		maxWeight = 1000
		x         = 0
		y         = 0
	)
	truckSuccess := assert.Nil(t, gameEnv.Create_camion(name, x, y, maxWeight, 2))
	assert.NotNil(t, gameEnv.Create_camion("name", 11, 10, 1000, 5), "out of bonds")
	assert.NotNil(t, gameEnv.Create_camion("name", x, y, 1000, 5), "cell already occupied")
	assert.NotNil(t, gameEnv.Create_camion("name", x, y, -1000, 5), "negative weight")
	assert.NotNil(t, gameEnv.Create_camion("name", x, y, 1000, -5), "negative turns")

	if truckSuccess {
		truck := gameEnv.Trucks[0]
		assert.Equal(t, 2, truck.Get_time_max())
		assert.Equal(t, name, truck.Get_name())
		assert.Equal(t, 0, truck.Get_current_weight())
		assert.Equal(t, maxWeight, truck.Get_max_weight())
		resx, resy := truck.Get_position()
		assert.Equal(t, resx, x)
		assert.Equal(t, resy, y)
		assert.Equal(t, game.CAMION, truck.Get_type())
		assert.Equal(t, "", truck.Get_status(), "no status yet")
		assert.Equal(t, true, truck.Is_present(), "truck hasn't moved yet")
		pack := game.Colis{}
		assert.Nil(t, truck.AddPackage(&pack), "should add the package")
		// can't check if package is added because Colis weights 0
		assert.Nil(t, truck.Move(0, 0, &gameEnv.Map))
		assert.Equal(t, "GONE", truck.Get_status(), "status GONE because it moved")
		assert.Equal(t, false, truck.Is_present(), "truck has moved")
		assert.Nil(t, truck.NextTurn())
		assert.Equal(t, false, truck.Is_present(), "truck has yet to return 1")
		assert.Equal(t, 1, truck.Get_comeback(), "comeback should equal 1")
		assert.Equal(t, "GONE", truck.Get_status(), "status GONE because it hasn't comeback")
		assert.Nil(t, truck.NextTurn())
		assert.Equal(t, false, truck.Is_present(), "truck has yet to return 2")
		assert.Equal(t, 0, truck.Get_comeback(), "comeback should equal 0")
		assert.NotNil(t, truck.Move(0, 0, &gameEnv.Map), "truck is GONE, shoudn't be able to move")
		assert.Nil(t, truck.NextTurn())
		assert.Equal(t, "", truck.Get_status(), "no status because new turn")
		assert.Equal(t, true, truck.Is_present(), "truck should be back")
		assert.Nil(t, truck.Wait())
		assert.NotNil(t, truck.Wait(), "can do only one action per turn")
		assert.Nil(t, truck.NextTurn())
		if assert.Equal(t, "", truck.Get_status()) {
			assert.NotNil(t, truck.NextTurn(), "should fail because no actions were taken this turn")
		}
	}
}
