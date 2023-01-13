package algo

import (
	"github.com/darkelle24/go-epitech/game"
)

func sendTruck(truck *game.Camion, gameEnv *game.Game) {
	if truck.Is_present() && truck.Get_current_weight() == 0 {
		_ = truck.Wait()
	}
	for _, transp := range gameEnv.Transps {
		var tool game.Tool = transp
		if truck.Get_distance(&tool) < truck.Get_time_max() && transp.Has_Colis() && truck.Get_max_weight()-truck.Get_current_weight() >= transp.Get_Colis().Get_current_weight() {
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
