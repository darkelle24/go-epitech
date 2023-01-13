package game

import (
	"math"
)

// Transpalette is a Tool that can pickup Colis and bring them to the Camion
type Transpalette struct {
	name       string
	status     string
	x          int
	y          int
	colis      *Colis
	lastDroped *Colis
} // (tool, movable)

// Move is used to move the Transpalette
func (trans *Transpalette) Move(x, y int, floor *[][]Floor) error {
	xlen := len(*floor)
	ylen := len((*floor)[0])
	newx := trans.x + x
	newy := trans.y + y
	if (x != 0 && y != 0) || x > 1 || x < -1 || y > 1 || y < -1 {
		return errTooFar
	}
	if (*floor)[newx][newy].Tool != nil {
		return errAlreadyOccupied
	}
	if (newx < 0 && newx >= xlen) || (newy < 0 && newy >= ylen) {
		return errOutOfBonds
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

// Wait for the Transpalette to skip this turn
func (trans *Transpalette) Wait() {
	trans.status = "WAIT"
}

// Take is used to take a Colis on a designed cell
func (trans *Transpalette) Take(x, y int, floor *[][]Floor) error {
	var tile = (*floor)[x][y]
	pack, ok := tile.Tool.(*Colis)
	if !ok || trans.colis != nil {
		return errWrongTarget
	}
	trans.colis = pack
	(*floor)[x][y].Tool = nil
	trans.status = "TAKE"
	return nil
}

// Drop is used to drop your Colis in the Truck
func (trans *Transpalette) Drop(x, y int, floor *[][]Floor) error {
	var tile = (*floor)[x][y]
	truck, ok := tile.Tool.(*Camion)
	if !ok || trans.colis == nil || !truck.Is_present() {
		return errWrongTarget
	}
	if err := truck.AddPackage(trans.colis); err != nil {
		return err
	}
	trans.colis.SetDelivered()
	trans.lastDroped = trans.colis
	trans.colis = nil
	trans.status = "LEAVE"
	return nil
}

// Get_name returns the Transpalette's name
func (trans *Transpalette) Get_name() string {
	return trans.name
}

// Get_type returns the Transpalette's type
func (trans *Transpalette) Get_type() TypeTool {
	return TRANSPALET
}

// Get_status returns the Transpalette's status
func (trans *Transpalette) Get_status() string {
	return trans.status
}

// Get_position returns the Transpalette's position x, y
func (trans *Transpalette) Get_position() (x int, y int) {
	return trans.x, trans.y
}

// Get_Colis returns the Transpalette's Colis
func (trans *Transpalette) Get_Colis() *Colis {
	return trans.colis
}

// Get_distance returns the Transpalette's distance to another tool
func (trans *Transpalette) Get_distance(ctool *Tool) int {
	tool := *ctool
	t_x, t_y := tool.Get_position()
	x := math.Abs(float64(trans.x) - float64(t_x))
	y := math.Abs(float64(trans.y) - float64(t_y))
	return int(x) + int(y) - 1
}

// Has_Colis returns true if Transpalette is carying a Colis
func (trans *Transpalette) Has_Colis() bool {
	return trans.colis != nil
}

// NextTurn processes actions taken this turn
func (trans *Transpalette) NextTurn() error {
	switch trans.status {
	case "GO":
		PrintTransMove(trans, trans.x, trans.y)
	case "WAIT":
		PrintTransWaiting(trans)
	case "TAKE":
		PrintTransPickup(trans, trans.colis)
	case "LEAVE":
		PrintTransDrop(trans, trans.lastDroped)
	default:
		return errNoAction
	}
	trans.status = ""
	return nil
}
