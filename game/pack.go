package game

import "math"

// Colis is the Tool that is picked up by Transpalette and droped in Camion
type Colis struct {
	name      string
	size      int
	x         int
	y         int
	delivered bool
} // (tool, weight)

// Get_name returns Colis's name
func (pack *Colis) Get_name() string {
	return pack.name
}

// Get_type returns Colis's type
func (pack *Colis) Get_type() TypeTool {
	return COLIS
}

// Get_position returns Colis's position
func (pack *Colis) Get_position() (x int, y int) {
	return pack.x, pack.y
}

// Get_distance returns Colis's distance to another tool
func (pack *Colis) Get_distance(ctool *Tool) int {
	tool := *ctool
	t_x, t_y := tool.Get_position()
	x := math.Abs(float64(pack.x) - float64(t_x))
	y := math.Abs(float64(pack.y) - float64(t_y))
	return int(x) + int(y) - 1
}

// Get_current_weight returns Colis's current weight
func (pack *Colis) Get_current_weight() int {
	return pack.size
}

// Get_max_weight returns Colis's max weight
func (pack *Colis) Get_max_weight() int {
	return pack.size
}

// SetDelivered sets Colis's delivered to true
func (pack *Colis) SetDelivered() {
	pack.delivered = true
}

// IsDelivered check and returns true if colis is delivered
func (pack *Colis) IsDelivered() bool {
	return pack.delivered
}
