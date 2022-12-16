package game

type TypeTool int

const (
	COLIS TypeTool = iota
	TRANSPALET
	CAMION
)

type tool interface {
	get_name() string
	get_type() TypeTool
}

type weight interface {
	get_current_weight() int
	get_max_weight() int
}

type movable interface {
	move(int, int, [][]floor)
	get_status() string
	get_position() (int, int)
}

type waiter interface {
	get_time_max() int
	get_comeback() int
	is_present() bool
}

// Floor
type floor struct {
	tool *tool
}

// OBjects
type camion struct {
	name         string
	poids_max    int
	turn_max     int
	turn_current int
	status       string
	x            int
	y            int
	colis_list   *colis
} //(tool, weight, movable, waiter)

type transpalette struct {
	name   string
	status string
	x      int
	y      int
	colis  *colis
} //(tool, movable)

type colis struct {
	name string
	size int
} //(tool, weight)

type game_functions interface {
	create_map(int, int)                      //x y
	create_transpallete(string, int, int)     // name, x, y
	create_colis(string, int, int, int)       // name, x, y, weight
	create_camion(string, int, int, int, int) // name, x, y, max_weight, turn_max
	next_turn()
}

type game struct {
	Map        [][]floor
	Turn       int
	GameObject []tool
} //(game_functions)
