package game

import (
	"fmt"
	"math"
)

// Camion is the Tool that goes to deliver the Colis and comes back after turn_max turns
type Camion struct {
	name           string
	poids_max      int
	weight_current int
	turn_max       int
	turn_current   int
	status         string
	x              int
	y              int
	colis_list     []*Colis
} // (tool, weight, movable, waiter)

// Camion methods

// Move makes the Camion go and be absent for turn_max turns
func (truck *Camion) Move(int, int, *[][]Floor) error {
	if !truck.Is_present() {
		return errWrongAction
	}
	const currentTurnOffset = 1
	truck.turn_current = truck.turn_max + currentTurnOffset
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

// Get_time_max returns Camion's turn_max
func (truck *Camion) Get_time_max() int {
	return truck.turn_max
}

// Get_comeback returns Camion's turn_current
func (truck *Camion) Get_comeback() int {
	return truck.turn_current
}

// Is_present returns true Camion's is present
func (truck *Camion) Is_present() bool {
	return truck.turn_current == 0
}

// NextTurn processes actions taken this turn
func (truck *Camion) NextTurn() error {
	if truck.turn_current > 0 {
		truck.turn_current--
		truck.status = "GONE"
		if truck.turn_current == 0 {
			truck.colis_list = []*Colis{}
		}
	}
	switch truck.status {
	case "WAITING":
		PrintTruckWaiting(truck)
	case "GONE":
		PrintTruckDepart(truck)
	default:
		return errNoAction
	}
	truck.status = ""
	return nil
}

// Get_name returns Camion's name
func (truck *Camion) Get_name() string {
	return truck.name
}

// Get_type returns Camion's type
func (truck *Camion) Get_type() TypeTool {
	return CAMION
}

// Get_current_weight returns Camion's weight
func (truck *Camion) Get_current_weight() int {
	return truck.weight_current
}

// Get_max_weight returns Camion's max weight
func (truck *Camion) Get_max_weight() int {
	return truck.poids_max
}

// Get_status returns Camion's status
func (truck *Camion) Get_status() string {
	return truck.status
}

// Get_position returns Camion's position x, y
func (truck *Camion) Get_position() (x int, y int) {
	return truck.x, truck.y
}

// Get_distance returns Camion's distance to an other tool
func (truck *Camion) Get_distance(ctool *Tool) int {
	tool := *ctool
	t_x, t_y := tool.Get_position()
	x := math.Abs(float64(truck.x) - float64(t_x))
	y := math.Abs(float64(truck.y) - float64(t_y))
	return int(x) + int(y) - 1
}

// AddPackage tries to add a package from a Transpalette to the Truck
func (truck *Camion) AddPackage(pack *Colis) error {
	totalWeight := truck.weight_current + pack.Get_current_weight()
	fmt.Println("---123---", totalWeight, truck.weight_current, pack.Get_current_weight(), truck.Get_max_weight())
	if totalWeight > truck.Get_max_weight() {
		return errWrongAction
	}
	truck.weight_current = totalWeight
	truck.colis_list = append(truck.colis_list, pack)
	return nil
}
