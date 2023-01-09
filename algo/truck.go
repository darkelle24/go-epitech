package algo

import (
	"github.com/darkelle24/go-epitech/game"
)

func SendTruck(truck *game.Camion, gameEnv *game.Game) {
	if truck.Is_present() && truck.Get_current_weight() == 0 {
		truck.Wait()
		return
	}
	for _, t := range gameEnv.Transps {
		var i game.Tool = t
		if truck.Get_distance(&i) <= truck.Get_time_max()-1 && t.Has_Colis() && truck.Get_max_weight()-truck.Get_current_weight() >= t.Get_Colis().Get_current_weight() {
			truck.Wait()
			return
		}
	}
	truck.Move(0, 0, &gameEnv.Map)
}
