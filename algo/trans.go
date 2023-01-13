package algo

import (
	"sort"

	"github.com/darkelle24/go-epitech/game"
)

// PackageManager is a type that help make orders and move each Colis
type PackageManager struct {
	Pack       *game.Colis
	Trans      []*game.Transpalette
	FinalTrans *game.Transpalette
	FinalTruck *game.Camion
}

func filterTrans(pack game.Tool, allTransp []*game.Transpalette) []*game.Transpalette {
	newTransp := make([]*game.Transpalette, 0)

	for _, trans := range allTransp {
		if trans.GetDistance(&pack) == allTransp[0].GetDistance(&pack) {
			newTransp = append(newTransp, trans)
		} else {
			return newTransp
		}
	}
	return newTransp
}

func checkIfTruckInSlice(toFind *game.Transpalette, trucks []*game.Transpalette) bool {
	for _, truck := range trucks {
		if truck == toFind {
			return true
		}
	}
	return false
}

func setFinalTransp(transps []*game.Transpalette, managers *[]*PackageManager) {
	for _, trans := range transps {
		var bestPack *PackageManager

		for _, pack := range *managers {
			if !checkIfTruckInSlice(trans, pack.Trans) {
				continue
			}
			if bestPack == nil && pack.FinalTrans == nil {
				bestPack = pack
				continue
			}
			var CurrentPack game.Tool = pack.Pack
			var BestPack game.Tool = bestPack.Pack
			println(bestPack.Trans)
			if pack.FinalTrans == nil && pack.FinalTruck.GetDistance(&CurrentPack) < bestPack.FinalTruck.GetDistance(&BestPack) {
				bestPack = pack
			}
		}
		bestPack.FinalTrans = trans
		bestPack.Trans = make([]*game.Transpalette, 0)
	}
}

// SetupManager makes the basic orders for the project
func SetupManager(gameEnv *game.Game) []*PackageManager {
	packsInfo := make([]*PackageManager, 0)

	for _, pack := range gameEnv.Packs {
		var tool game.Tool = pack
		trans := gameEnv.Transps
		trucks := gameEnv.Trucks

		sort.Slice(trans, func(i, j int) bool {
			return trans[i].GetDistance(&tool) < trans[j].GetDistance(&tool)
		})
		sort.Slice(trucks, func(i, j int) bool {
			return trucks[i].GetDistance(&tool) < trucks[j].GetDistance(&tool)
		})
		packsInfo = append(packsInfo, &PackageManager{Pack: pack, Trans: filterTrans(tool, trans), FinalTruck: trucks[0]})
	}
	setFinalTransp(gameEnv.Transps, &packsInfo)
	return packsInfo
}

func wichTransAvailable(transps []*game.Transpalette, managers []*PackageManager) *game.Transpalette {
	return nil
}

func moveToBox(trans *game.Transpalette, pack *game.Colis, ground *[][]game.Floor) {
	var toolTrans game.Tool = trans

	if pack.GetDistance(&toolTrans) == 0 {
		x, y := pack.GetPosition()
		_ = trans.Take(x, y, ground)
	}
}

func moveToTruck(trans *game.Transpalette, truck *game.Camion, ground *[][]game.Floor) {
	var toolTrans game.Tool = trans

	if truck.GetDistance(&toolTrans) == 0 && trans.GetColis() != nil {
		x, y := truck.GetPosition()
		_ = trans.Drop(x, y, ground)
	}
}

// UpdateManager handles all the manager for each Package and what actions should Transpalette be taking
func UpdateManager(gameEnv *game.Game, managers []*PackageManager) []*PackageManager {
	for i, manager := range managers {
		if manager.Pack.IsDelivered() {
			managers = append(managers[:i], managers[i+1:]...)
			continue
		}
		available := wichTransAvailable(gameEnv.Transps, managers)
		if manager.FinalTrans == nil && available == nil {
			continue
		} else if manager.FinalTrans == nil {
			manager.FinalTrans = available
		}
		if manager.FinalTrans.GetColis() == nil {
			moveToBox(manager.FinalTrans, manager.Pack, &gameEnv.Map)
		} else {
			moveToTruck(manager.FinalTrans, manager.FinalTruck, &gameEnv.Map)
		}
	}
	return managers
}
