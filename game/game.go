package game

import (
	"errors"
	"fmt"
)

type TypeTool int

const (
	COLIS TypeTool = iota
	TRANSPALET
	CAMION
)

type Tool interface {
	Get_name() string
	Get_type() TypeTool
	Get_position() (int, int)
}

type Weight interface {
	Get_current_weight() int
	Get_max_weight() int
}

type Movable interface {
	Move(int, int, *[][]Floor) error
	Get_status() string
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
	colis_list   []*Colis
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
	x    int
	y    int
} //(tool, weight)

type game_functions interface {
	Set_turns(int)                                  // x y
	Get_turns() int                                 // x y
	Create_map(int, int) error                      // x y
	Create_transpallete(string, int, int) error     // name, x, y
	Create_colis(string, int, int, int) error       // name, x, y, weight
	Create_camion(string, int, int, int, int) error // name, x, y, max_weight, turn_max
	Next_turn()
}

type Game struct {
	Map       [][]Floor
	Turn      int
	turns     int
	ToolsList []Tool
} //(game_functions)

// Game methods

func (game *Game) Set_turns(turns int) {
	game.turns = turns
}

func (game *Game) Get_turns() int {
	return game.turns
}

func (game *Game) Create_map(x, y int) error {
	if game.Map != nil {
		return errors.New("map is already created")
	}
	if x <= 0 || y <= 0 {
		return errors.New("map can't be of negative length")
	}
	game.Map = make([][]Floor, x)
	for i := range game.Map {
		game.Map[i] = make([]Floor, y)
	}
	return nil
}

func checkToolCreation(game *Game, x, y int) error {
	xlen := len(game.Map)
	ylen := len(game.Map[0])
	if game.Map[x][y].Tool != nil {
		return errors.New("can't create here this cell is already occupied")
	}
	if (x < 0 && x >= xlen) || (y < 0 && y >= ylen) {
		return errors.New("tool must be created within the map")
	}
	return nil
}

func (game *Game) Create_transpallete(name string, x, y int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	var trans Tool = &Transpalette{name: name, x: x, y: y}
	game.ToolsList = append(game.ToolsList, trans)
	tile := Floor{Tool: trans}
	game.Map[x][y] = tile
	return nil
}

func (game *Game) Create_colis(name string, x, y, weight int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	if weight <= 0 {
		return errors.New("colis can't be created with a negative weight")
	}
	var pack Tool = &Colis{name: name, x: x, y: y, size: weight}
	game.ToolsList = append(game.ToolsList, pack)
	tile := Floor{Tool: pack}
	game.Map[x][y] = tile
	return nil
}

func (game *Game) Create_camion(name string, x, y, max_weight, turn_max int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	if max_weight <= 0 {
		return errors.New("camion can't be created with a negative max weight")
	}
	if turn_max <= 0 {
		return errors.New("camion can't be created with a negative turn max")
	}
	var truck Tool = &Camion{name: name, x: x, y: y, poids_max: max_weight, turn_max: turn_max}
	game.ToolsList = append(game.ToolsList, truck)
	tile := Floor{Tool: truck}
	game.Map[x][y] = tile
	return nil
}

func (game *Game) Next_turn() {
	// apelle next_turn sur tout les waiter?
	game.Turn += 1
	fmt.Println("Tour", game.Turn)
	// for _, v := range game.Map {
	// 	fmt.Println(v)
	// }
	// fmt.Println(game.ToolsList)
	// time.Sleep(time.Second)
}

// Camion methods
func (truck *Camion) Move(int, int, *[][]Floor) error {
	if !truck.Is_present() {
		return errors.New("truck can't move because it is GONE")
	}
	truck.turn_current = truck.turn_max
	truck.status = "GONE"
	return nil
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
		if truck.turn_current == 0 {
			truck.colis_list = []*Colis{}
		}
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

func (truck *Camion) AddPackage(pack *Colis) {
	// error if > max weight
	truck.colis_list = append(truck.colis_list, pack)
}

// Transpalette methods

func (trans *Transpalette) Move(x, y int, floor *[][]Floor) error {
	xlen := len(*floor)
	ylen := len((*floor)[0])
	newx := trans.x + x
	newy := trans.y + y
	if (x != 0 && y != 0) || x > 1 || x < -1 || y > 1 || y < -1 {
		return errors.New("can't move more than one cell")
	}
	if (*floor)[newx][newy].Tool != nil {
		return errors.New("can't move on an occupied cell")
	}
	if (newx < 0 && newx >= xlen) || (newy < 0 && newy >= ylen) {
		return errors.New("cannot move out of the map")
	}
	if (*floor)[newx][trans.y+y] == (Floor{}) {
		(*floor)[newx][trans.y+y].Tool = trans
		(*floor)[trans.x][trans.y].Tool = nil
		trans.x = newx
		trans.y = newy
	}
	trans.status = "GO"
	return nil
}

func (trans *Transpalette) Wait() {
	trans.status = "WAIT"
}

func (trans *Transpalette) Take(x, y int, floor *[][]Floor) error {
	var tile = (*floor)[x][y]
	pack, ok := tile.Tool.(*Colis)
	if !ok || trans.colis != nil {
		return errors.New("can't pickup package on this tile")
	}
	trans.colis = pack
	(*floor)[x][y].Tool = nil
	trans.status = "TAKE"
	return nil
}

func (trans *Transpalette) Drop(x, y int, floor *[][]Floor) error {
	var tile = (*floor)[x][y]
	truck, ok := tile.Tool.(*Camion)
	if !ok || trans.colis == nil || !truck.Is_present() {
		return errors.New("can't drop package on this tile")
	}
	truck.AddPackage(trans.colis)
	trans.colis = nil
	trans.status = "LEAVE"
	return nil
}

func (trans *Transpalette) Get_name() string {
	return trans.name
}

func (trans *Transpalette) Get_type() TypeTool {
	return TRANSPALET
}

func (trans *Transpalette) Get_status() string {
	return trans.status
}

func (trans *Transpalette) Get_position() (int, int) {
	return trans.x, trans.y
}

// Colis methods

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
