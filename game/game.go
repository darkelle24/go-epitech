// Package that implements the core functions of the game
// package game is used to initialize the game and provide a state of it
package game

import (
	"errors"
	"fmt"
)

// TypeTool is the type for the Tool's type
type TypeTool int

// COLIS TRANSPALET and CAMION are the 3 types of Tools
const (
	COLIS TypeTool = iota
	TRANSPALET
	CAMION
)

// Tool is an interface for Camion, Colis and Transpalette
type Tool interface {
	Get_name() string
	Get_type() TypeTool
	Get_position() (int, int)
	Get_distance(*Tool) int
}

// Weight is an interface for Colis and Camion
type Weight interface {
	Get_current_weight() int
	Get_max_weight() int
}

// Movable is an interface for Camion and Transpalette
type Movable interface {
	Move(int, int, *[][]Floor) error
	Get_status() string
	NextTurn() error
}

// Waiter is an interface for Camion
type Waiter interface {
	Get_time_max() int
	Get_comeback() int
	Is_present() bool
}

// Floor is the Map
type Floor struct {
	Tool Tool
}

// Game is the state of the game
type Game struct {
	Map       [][]Floor
	Turn      int
	turns     int
	ToolsList []Tool
	Trucks    []*Camion
	Packs     []*Colis
	Transps   []*Transpalette
} // (game_functions)

// Errors

var errAlreadyCreated = errors.New("already created")
var errNegativeValue = errors.New("tried to provide negative value for a positive field")
var errMapNotCreated = errors.New("must have a map to create a tool")
var errOutOfBonds = errors.New("targeted cell is out of bonds")
var errAlreadyOccupied = errors.New("cell is already occupied")
var errTooFar = errors.New("cannot target this far")
var errWrongTarget = errors.New("cannot target this")
var errNoAction = errors.New("no action was done last turn")
var errWrongAction = errors.New("cannot do this action")

// Game methods

// Set_turns sets the total number of turns in the game
func (game *Game) Set_turns(turns int) {
	game.turns = turns
}

// Get_turns get the total number of turns in the game
func (game *Game) Get_turns() int {
	return game.turns
}

// Create_map initialize the map by with the dimensions in the arguments
func (game *Game) Create_map(x, y int) error {
	if game.Map != nil {
		return errAlreadyCreated
	}
	if x <= 0 || y <= 0 {
		return errNegativeValue
	}
	game.Map = make([][]Floor, x)
	for i := range game.Map {
		game.Map[i] = make([]Floor, y)
	}
	return nil
}

func checkToolCreation(game *Game, x, y int) error {
	if game.Map == nil {
		return errMapNotCreated
	}
	xlen := len(game.Map)
	ylen := len(game.Map[0])
	if x < 0 || x >= xlen || y < 0 || y >= ylen {
		return errOutOfBonds
	}
	if game.Map[x][y].Tool != nil {
		return errAlreadyOccupied
	}
	return nil
}

// Create_transpallete creates a transpallete Tool
func (game *Game) Create_transpallete(name string, x, y int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	trans := &Transpalette{name: name, x: x, y: y}
	game.ToolsList = append(game.ToolsList, trans)
	game.Transps = append(game.Transps, trans)
	tile := Floor{Tool: trans}
	game.Map[x][y] = tile
	return nil
}

// Create_colis creates a colis Tool
func (game *Game) Create_colis(name string, x, y, weight int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	if weight <= 0 {
		return errNegativeValue
	}
	pack := &Colis{name: name, x: x, y: y, size: weight}
	game.Packs = append(game.Packs, pack)
	game.ToolsList = append(game.ToolsList, pack)
	tile := Floor{Tool: pack}
	game.Map[x][y] = tile
	return nil
}

// Create_camion creates a camion Tool
func (game *Game) Create_camion(name string, x, y, max_weight, turn_max int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	if max_weight <= 0 {
		return errNegativeValue
	}
	if turn_max <= 0 {
		return errNegativeValue
	}
	truck := Camion{name: name, x: x, y: y, poids_max: max_weight, turn_max: turn_max}
	game.ToolsList = append(game.ToolsList, &truck)
	game.Trucks = append(game.Trucks, &truck)
	tile := Floor{Tool: &truck}
	game.Map[x][y] = tile
	return nil
}

/* NextTurn calls the end of turn logic
** icrementing the current turn and calling the Movable NextTurn functions
 */
func (game *Game) NextTurn() {
	game.Turn++
	PrintNextTurn(game.Turn)
	for _, v := range game.Transps {
		err := v.NextTurn()
		if err != nil {
			PrintError(err)
		}
	}
	for _, v := range game.Trucks {
		err := v.NextTurn()
		if err != nil {
			PrintError(err)
		}
	}
	fmt.Println("")
}

// IsAllDelivered checks if all colis are delivered
func (game *Game) IsAllDelivered() bool {
	for _, box := range game.Packs {
		if !box.IsDelivered() {
			return false
		}
	}
	return true
}

// IsDone checks if the game is over
func (game *Game) IsDone() bool {
	if game.Turn >= game.Get_turns() {
		return true
	}
	if game.IsAllDelivered() {
		return true
	}
	return false
}

// EndStateCharacter returns 😎 if every colis is delivered else 🙂
func (game *Game) EndStateCharacter() string {
	if game.IsAllDelivered() {
		return "😎"
	}
	return "🙂"
}

// PrintMap prints the map of the current state
func (game *Game) PrintMap() {
	fmt.Println("Map:")
	for i := range game.Map[0] {
		for _, line := range game.Map {
			if line[i].Tool != nil {
				fmt.Printf("%s\t", line[i].Tool.Get_name())
			} else {
				fmt.Printf("--------\t")
			}
		}
		fmt.Println()
	}
}

// PrintTools print all tools
func (game *Game) PrintTools() {
	fmt.Println("Tools:")
	for _, elm := range game.ToolsList {
		fmt.Printf("%s\t", elm.Get_name())
	}
	fmt.Println("")
}
