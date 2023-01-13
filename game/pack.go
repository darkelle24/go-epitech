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

// GetName returns Colis's name
func (pack *Colis) GetName() string {
	return pack.name
}

// GetType returns Colis's type
func (pack *Colis) GetType() TypeTool {
	return COLIS
}

// GetPosition returns Colis's position
func (pack *Colis) GetPosition() (x int, y int) {
	return pack.x, pack.y
}

// GetDistance returns Colis's distance to another tool
func (pack *Colis) GetDistance(ctool *Tool) int {
	tool := *ctool
	tX, tY := tool.GetPosition()
	x := math.Abs(float64(pack.x) - float64(tX))
	y := math.Abs(float64(pack.y) - float64(tY))
	return int(x) + int(y) - 1
}

// GetCurrentWeight returns Colis's current weight
func (pack *Colis) GetCurrentWeight() int {
	return pack.size
}

// GetMaxWeight returns Colis's max weight
func (pack *Colis) GetMaxWeight() int {
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
