package game

import (
	"errors"
	"math"
)

type Transpalette struct {
	name       string
	status     string
	x          int
	y          int
	colis      *Colis
	lastDroped *Colis
} //(tool, movable)

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
	if err := truck.AddPackage(trans.colis); err != nil {
		return err
	}
	trans.colis.SetDelivered()
	trans.lastDroped = trans.colis
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

func (trans *Transpalette) Get_Colis() *Colis {
	return trans.colis
}

func (trans *Transpalette) Get_distance(ctool *Tool) int {
	tool := *ctool
	t_x, t_y := tool.Get_position()
	x := math.Abs(float64(trans.x) - float64(t_x))
	y := math.Abs(float64(trans.y) - float64(t_y))
	return int(x) + int(y) - 1
}

func (trans *Transpalette) Has_Colis() bool {
	return trans.colis != nil
}

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
		return errors.New("no action was done last turn")
	}
	trans.status = ""
	return nil
}
