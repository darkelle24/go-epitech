package game

type TypeTool int

const (
	COLIS TypeTool = iota
	TRANSPALET
	CAMION
)

type Tool interface {
	Get_name() string
	Get_type() TypeTool
}

type Weight interface {
	Get_current_weight() int
	Get_max_weight() int
}

type Movable interface {
	Move(int, int, [][]Floor)
	Get_status() string
	Get_position() (int, int)
}

type Waiter interface {
	Get_time_max() int
	Get_comeback() int
	Is_present() bool
	Next_turn()
}

// Floor
type Floor struct {
	Tool Tool
}

// OBjects
type Camion struct {
	name         string
	poids_max    int
	turn_max     int
	turn_current int
	status       string
	x            int
	y            int
	colis_list   *Colis
} //(tool, weight, movable, waiter)

type Transpalette struct {
	name   string
	status string
	x      int
	y      int
	colis  *Colis
} //(tool, movable)

type Colis struct {
	name string
	size int
} //(tool, weight)

type game_functions interface {
	Create_map(int, int)                      // x y
	Create_transpallete(string, int, int)     // name, x, y
	Create_colis(string, int, int, int)       // name, x, y, weight
	Create_camion(string, int, int, int, int) // name, x, y, max_weight, turn_max
	Next_turn()                               // apelle next_turn sur tout les waiter?
}

type Game struct {
	Map       [][]Floor
	Turn      int
	ToolsList []Tool
} //(game_functions)

// Game methods
func (game *Game) Create_map(x, y int) {
	game.Map = make([][]Floor, x)
	for i := range game.Map {
		game.Map[i] = make([]Floor, y)
	}
}

func (game *Game) Create_transpallete(name string, x, y int) {
}

func (game *Game) Create_colis(name string, x, y, weight int) {

}

func (game *Game) Create_camion(name string, x, y, max_weight, turn_max int) {
	var truck Tool
	truck = &Camion{name: name, x: x, y: y, poids_max: max_weight, turn_max: turn_max}
	game.ToolsList = append(game.ToolsList, truck)
	cell := Floor{Tool: truck}
	game.Map[x][y] = cell
}

// Camion methods
func (truck *Camion) Move(int, int, [][]Floor) {
	truck.turn_current = truck.turn_max
	truck.status = "GONE"
}

func (truck *Camion) Wait() {
	truck.status = "WAITING"
}

func (truck *Camion) Get_time_max() int {
	return truck.turn_max
}

func (truck *Camion) Get_comeback() int {
	return truck.turn_current
}

func (truck *Camion) Is_present() bool {
	return truck.turn_current == 0
}

func (truck *Camion) Next_turn() {
	if truck.turn_current > 0 {
		truck.turn_current -= 1
	}
}

func (truck *Camion) Get_name() string {
	return truck.name
}

func (truck *Camion) Get_type() TypeTool {
	return CAMION
}

func (truck *Camion) Get_current_weight() int {
	// add all weight from truck.colis_list
	return 0
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

//
