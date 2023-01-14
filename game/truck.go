package game

import (
	"math"
)

// Camion is the Tool that goes to deliver the Colis and comes back after turnMax turns
type Camion struct {
	name          string
	poidsMax      int
	weightCurrent int
	turnMax       int
	turnCurrent   int
	status        string
	x             int
	y             int
	colisList     []*Colis
} // (tool, weight, movable, waiter)

// Camion methods

// Move makes the Camion go and be absent for turnMax turns
func (truck *Camion) Move(int, int, *[][]Floor) error {
	if !truck.IsPresent() {
		return errWrongAction
	}
	const currentTurnOffset = 1
	truck.turnCurrent = truck.turnMax + currentTurnOffset
	truck.status = "GONE"
	return nil
}

// Wait makes the camion skip this turn
func (truck *Camion) Wait() error {
	if truck.status != "" {
		return errWrongAction
	}
	truck.status = "WAITING"
	return nil
}

// GetTimeMax returns Camion's turnMax
func (truck *Camion) GetTimeMax() int {
	return truck.turnMax
}

// GetComeback returns Camion's turnCurrent
func (truck *Camion) GetComeback() int {
	return truck.turnCurrent
}

// IsPresent returns true Camion's is present
func (truck *Camion) IsPresent() bool {
	return truck.turnCurrent == 0
}

// NextTurn processes actions taken this turn
func (truck *Camion) NextTurn() error {
	if truck.turnCurrent > 0 {
		truck.turnCurrent--
		truck.status = "GONE"
	}
	switch truck.status {
	case "WAITING":
		PrintTruckWaiting(truck)
	case "GONE":
		PrintTruckDepart(truck)
		if truck.turnCurrent == 0 {
			truck.colisList = []*Colis{}
			truck.weightCurrent = 0
		}
	default:
		return errNoAction
	}
	truck.status = ""
	return nil
}

// GetName returns Camion's name
func (truck *Camion) GetName() string {
	return truck.name
}

// GetType returns Camion's type
func (truck *Camion) GetType() TypeTool {
	return CAMION
}

// GetCurrentWeight returns Camion's weight
func (truck *Camion) GetCurrentWeight() int {
	return truck.weightCurrent
}

// GetMaxWeight returns Camion's max weight
func (truck *Camion) GetMaxWeight() int {
	return truck.poidsMax
}

// GetStatus returns Camion's status
func (truck *Camion) GetStatus() string {
	return truck.status
}

// GetPosition returns Camion's position x, y
func (truck *Camion) GetPosition() (x int, y int) {
	return truck.x, truck.y
}

// GetDistance returns Camion's distance to an other tool
func (truck *Camion) GetDistance(ctool *Tool) int {
	tool := *ctool
	tX, tY := tool.GetPosition()
	x := math.Abs(float64(truck.x) - float64(tX))
	y := math.Abs(float64(truck.y) - float64(tY))
	return int(x) + int(y) - 1
}

// AddPackage tries to add a package from a Transpalette to the Truck
func (truck *Camion) AddPackage(pack *Colis) error {
	totalWeight := truck.weightCurrent + pack.GetCurrentWeight()
	if totalWeight > truck.GetMaxWeight() {
		return errWrongAction
	}
	truck.weightCurrent = totalWeight
	truck.colisList = append(truck.colisList, pack)
	return nil
}
