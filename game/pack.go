package game

import "math"

type Colis struct {
	name      string
	size      int
	x         int
	y         int
	delivered bool
} //(tool, weight)

func (pack *Colis) Get_name() string {
	return pack.name
}

func (pack *Colis) Get_type() TypeTool {
	return COLIS
}

func (pack *Colis) Get_position() (int, int) {
	return pack.x, pack.y
}

func (pack *Colis) Get_distance(ctool *Tool) int {
	tool := *ctool
	t_x, t_y := tool.Get_position()
	x := math.Abs(float64(pack.x) - float64(t_x))
	y := math.Abs(float64(pack.y) - float64(t_y))
	return int(x) + int(y) - 1
}

func (pack *Colis) Get_current_weight() int {
	return pack.size
}

func (pack *Colis) Get_max_weight() int {
	return pack.size
}

func (pack *Colis) SetDelivered() {
	pack.delivered = true
}

func (pack *Colis) IsDelivered() bool {
	return pack.delivered
}
