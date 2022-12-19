package game

import "errors"

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
} //(tool, weight, movable, waiter)

// Camion methods
func (truck *Camion) Move(int, int, *[][]Floor) error {
	if !truck.Is_present() {
		return errors.New("truck can't move because it is GONE")
	}
	truck.turn_current = truck.turn_max
	truck.status = "GONE"
	PrintTruckDepart(truck)
	return nil
}

func (truck *Camion) Wait() error {
	if truck.status != "" {
		return errors.New("can only do one action a turn")
	}
	PrintTruckWaiting(truck)
	truck.status = "WAITING"
	return nil
}

func (truck *Camion) Get_time_max() int {
	return truck.turn_max
}

func (truck *Camion) Get_comeback() int {
	return truck.turn_current
}

func (truck *Camion) Is_present() bool {
	return truck.status != "GONE"
}

func (truck *Camion) NextTurn() error {
	if truck.status == "" {
		return errors.New("no action was done last turn")
	}
	truck.status = ""
	if truck.turn_current > 0 {
		truck.turn_current -= 1
		truck.status = "GONE"
		PrintTruckGone(truck)
		if truck.turn_current == 0 {
			truck.colis_list = []*Colis{}
		}
	}
	return nil
}

func (truck *Camion) Get_name() string {
	return truck.name
}

func (truck *Camion) Get_type() TypeTool {
	return CAMION
}

func (truck *Camion) Get_current_weight() int {
	return truck.weight_current
}

func (truck *Camion) Get_max_weight() int {
	return truck.poids_max
}

func (truck *Camion) Get_status() string {
	return truck.status
}

func (truck *Camion) Get_position() (int, int) {
	return truck.x, truck.y
}

func (truck *Camion) AddPackage(pack *Colis) error {
	totalWeight := truck.weight_current + pack.Get_current_weight()
	if totalWeight > truck.Get_max_weight() {
		return errors.New("adding this package would surchage the truck")
	}
	truck.weight_current = totalWeight
	truck.colis_list = append(truck.colis_list, pack)
	return nil
}
