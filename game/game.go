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
	Get_distance(*Tool) int
}

type Weight interface {
	Get_current_weight() int
	Get_max_weight() int
}

type Movable interface {
	Move(int, int, *[][]Floor) error
	Get_status() string
	NextTurn() error
}

type Waiter interface {
	Get_time_max() int
	Get_comeback() int
	Is_present() bool
}

type Floor struct {
	Tool Tool
}

type Game struct {
	Map       [][]Floor
	Turn      int
	turns     int
	ToolsList []Tool
	Trucks    []*Camion
	Packs     []*Colis
	Transps   []*Transpalette
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
	if game.Map == nil {
		return errors.New("must have a map to create a tool")
	}
	xlen := len(game.Map)
	ylen := len(game.Map[0])
	if x < 0 || x >= xlen || y < 0 || y >= ylen {
		return errors.New("tool must be created within the map")
	}
	if game.Map[x][y].Tool != nil {
		return errors.New("can't create here this cell is already occupied")
	}
	return nil
}

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

func (game *Game) Create_colis(name string, x, y, weight int) error {
	if err := checkToolCreation(game, x, y); err != nil {
		return err
	}
	if weight <= 0 {
		return errors.New("colis can't be created with a negative weight")
	}
	pack := &Colis{name: name, x: x, y: y, size: weight}
	game.Packs = append(game.Packs, pack)
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
	truck := Camion{name: name, x: x, y: y, poids_max: max_weight, turn_max: turn_max}
	game.ToolsList = append(game.ToolsList, &truck)
	game.Trucks = append(game.Trucks, &truck)
	tile := Floor{Tool: &truck}
	game.Map[x][y] = tile
	return nil
}

func (game *Game) NextTurn() {
	game.Turn += 1
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

func (game *Game) IsAllDelivered() bool {
	for _, box := range game.Packs {
		if !box.IsDelivered() {
			return false
		}
	}
	return true
}

func (game *Game) IsDone() bool {
	if game.Turn >= game.Get_turns() {
		return true
	}
	if game.IsAllDelivered() {
		return true
	}
	return false
}

func (game *Game) EndStateCharacter() string {
	if game.IsAllDelivered() {
		return "😎"
	}
	return "🙂"
}

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

func (game *Game) PrintTools() {
	fmt.Println("Tools:")
	for _, elm := range game.ToolsList {
		fmt.Printf("%s\t", elm.Get_name())
	}
	fmt.Println("")
}
