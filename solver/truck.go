package solver

import (
	"github.com/darkelle24/go-epitech/game"
)

func sendTruck(truck *game.Camion, gameEnv *game.Game) {
	if truck.IsPresent() && truck.GetCurrentWeight() == 0 {
		_ = truck.Wait()
	}
	for _, transp := range gameEnv.Transps {
		var tool game.Tool = transp
		if truck.GetDistance(&tool) < truck.GetTimeMax() && transp.HasColis() && truck.GetMaxWeight()-truck.GetCurrentWeight() >= transp.GetColis().GetCurrentWeight() {
			_ = truck.Wait()
			return
		}
	}
	_ = truck.Move(0, 0, &gameEnv.Map)
}

// UpdateTrucks is the update function that tells what they should do
func UpdateTrucks(gameEnv *game.Game) error {
	for _, truck := range gameEnv.Trucks {
		sendTruck(truck, gameEnv)
	}
	return nil
}
