package game

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

func (pack *Colis) Get_current_weight() int {
	return pack.size
}

func (pack *Colis) Get_max_weight() int {
	return pack.size
}

func (pack *Colis) SetDelivered() {
	pack.delivered = true
}
