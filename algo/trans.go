package algo

import (
	"math"
	"sort"

	"github.com/darkelle24/go-epitech/game"
)

// Point is a type that allow to plan ahead
type Point struct {
	X int
	Y int
}

// PackageManager is a type that help make orders and move each Colis
type PackageManager struct {
	Pack       *game.Colis
	Trans      []*game.Transpalette
	FinalTrans *game.Transpalette
	FinalTruck *game.Camion
	GoToPoints []Point
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

func isTransInSlice(toFind *game.Transpalette, transps []*game.Transpalette) bool {
	for _, tran := range transps {
		if toFind == tran {
			return true
		}
	}
	return false
}

func wichTransAvailable(transps []*game.Transpalette, managers []*PackageManager) *game.Transpalette {
	useds := make([]*game.Transpalette, 0)

	for _, man := range managers {
		if man.FinalTrans != nil {
			useds = append(useds, man.FinalTrans)
		}
	}
	for _, tran := range transps {
		if !isTransInSlice(tran, useds) {
			return tran
		}
	}
	return nil
}

func moveToBox(trans *game.Transpalette, pack *game.Colis, points []Point, ground *[][]game.Floor) []Point {
	var toolTrans game.Tool = trans
	newPoint := make([]Point, 0)

	if pack.GetDistance(&toolTrans) == 0 {
		x, y := pack.GetPosition()
		_ = trans.Take(x, y, ground)
	} else if len(points) > 0 {
		_ = trans.Move(points[0].X, points[0].Y, ground)
		return points[1:]
	} else {
		tx, ty := trans.GetPosition()
		cx, cy := pack.GetPosition()
		var nx, ny int
		if math.Abs(float64(tx)-float64(cx)) != 0 {
			if tx-cx > 0 {
				nx = -1
			} else {
				nx = 1
			}
		}
		if math.Abs(float64(ty)-float64(cy)) != 0 {
			if ty-cy > 0 {
				ny = -1
			} else {
				ny = 1
			}
		}
		if nx != 0 && (*ground)[tx+nx][ty].Tool == nil {
			_ = trans.Move(nx, 0, ground)
		} else if ny != 0 && (*ground)[tx][ty+ny].Tool == nil {
			_ = trans.Move(0, ny, ground)
		} else {
			if nx == 0 {
				newPoint = append(newPoint, Point{X: 1, Y: 0})
				newPoint = append(newPoint, Point{X: 0, Y: ny})
			} else if ny == 0 {
				newPoint = append(newPoint, Point{X: 0, Y: 1})
				newPoint = append(newPoint, Point{X: nx, Y: 0})
			}
		}
	}
	return newPoint
}

func moveToTruck(trans *game.Transpalette, truck *game.Camion, points []Point, ground *[][]game.Floor) []Point {
	var toolTrans game.Tool = trans
	newPoint := make([]Point, 0)

	if truck.GetDistance(&toolTrans) == 0 && trans.GetColis() != nil {
		if !truck.IsPresent() {
			return newPoint
		}
		x, y := truck.GetPosition()
		_ = trans.Drop(x, y, ground)
	} else if len(points) > 0 {
		_ = trans.Move(points[0].X, points[0].Y, ground)
		return points[1:]
	} else {
		tx, ty := trans.GetPosition()
		cx, cy := truck.GetPosition()
		var nx, ny int
		if math.Abs(float64(tx)-float64(cx)) != 0 {
			if tx-cx > 0 {
				nx = -1
			} else {
				nx = 1
			}
		}
		if math.Abs(float64(ty)-float64(cy)) != 0 {
			if ty-cy > 0 {
				ny = -1
			} else {
				ny = 1
			}
		}
		if nx != 0 && (*ground)[tx+nx][ty].Tool == nil {
			_ = trans.Move(nx, 0, ground)
		} else if ny != 0 && (*ground)[tx][ty+ny].Tool == nil {
			_ = trans.Move(0, ny, ground)
		} else {
			newPoint := make([]Point, 0)
			if nx == 0 {
				newPoint = append(newPoint, Point{X: 1, Y: 0})
				newPoint = append(newPoint, Point{X: 0, Y: ny})
			} else if ny == 0 {
				newPoint = append(newPoint, Point{X: 0, Y: 1})
				newPoint = append(newPoint, Point{X: nx, Y: 0})
			}
		}
	}
	return newPoint
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
			manager.GoToPoints = moveToBox(manager.FinalTrans, manager.Pack, manager.GoToPoints, &gameEnv.Map)
		} else {
			manager.GoToPoints = moveToTruck(manager.FinalTrans, manager.FinalTruck, manager.GoToPoints, &gameEnv.Map)
		}
	}
	return managers
}
