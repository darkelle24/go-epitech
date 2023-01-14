// Package game is used to initialize the game and provide a state of it
// it implements the core functions of the game
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
	GetName() string
	GetType() TypeTool
	GetPosition() (int, int)
	GetDistance(*Tool) int
}

// Weight is an interface for Colis and Camion
type Weight interface {
	GetCurrentWeight() int
	GetMaxWeight() int
}

// Movable is an interface for Camion and Transpalette
type Movable interface {
	Move(int, int, *[][]Floor) error
	GetStatus() string
	NextTurn() error
}

// Waiter is an interface for Camion
type Waiter interface {
	GetTimeMax() int
	GetComeback() int
	IsPresent() bool
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

var (
	errAlreadyCreated  = errors.New("already created")
	errNegativeValue   = errors.New("tried to provide negative value for a positive field")
	errMapNotCreated   = errors.New("must have a map to create a tool")
	errOutOfBonds      = errors.New("targeted cell is out of bonds")
	errAlreadyOccupied = errors.New("cell is already occupied")
	errTooFar          = errors.New("cannot target this far")
	errWrongTarget     = errors.New("cannot target this")
	errNoAction        = errors.New("no action was done last turn")
	errWrongAction     = errors.New("cannot do this action")
)

// Game methods

// SetTurns sets the total number of turns in the game
func (game *Game) SetTurns(turns int) {
	game.turns = turns
}

// GetTurns get the total number of turns in the game
func (game *Game) GetTurns() int {
	return game.turns
}

// CreateMap initialize the map by with the dimensions in the arguments
func (game *Game) CreateMap(x, y int) error {
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

// CreateTranspallete creates a transpallete Tool
func (game *Game) CreateTranspallete(name string, x, y int) error {
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

// CreateColis creates a colis Tool
func (game *Game) CreateColis(name string, x, y, weight int) error {
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

// CreateCamion creates a camion Tool
func (game *Game) CreateCamion(name string, x, y, maxWeight, turnMax int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	if maxWeight <= 0 {
		return errNegativeValue
	}
	if turnMax <= 0 {
		return errNegativeValue
	}
	truck := Camion{name: name, x: x, y: y, poidsMax: maxWeight, turnMax: turnMax}
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
	if game.Turn >= game.GetTurns() {
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
				fmt.Printf("%s\t", line[i].Tool.GetName())
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
		fmt.Printf("%s\t", elm.GetName())
	}
	fmt.Println("")
}
