package game

import (
	"fmt"
	"time"
)

var colors = map[int]string{
	100: "YELLOW",
	200: "GREEN",
	500: "BLUE",
}

func UpdateMap(game *Game) {
	for _, v := range game.Map {
		fmt.Println(v)
	}
	fmt.Println(game.ToolsList)
	time.Sleep(time.Second)
}

func PrintNextTurn(turn int) {
	fmt.Println("\ntour", turn)
}

func PrintTransMove(trans *Transpalette, x, y int) {
	fmt.Printf("%v GO [%v,%v]\n", trans.Get_name(), x, y)
}

func PrintTruckDepart(truck *Camion) {
	fmt.Printf("%v GONE %v/%v\n", truck.Get_name(), truck.Get_current_weight(), truck.Get_max_weight())
}

func PrintTruckGone(truck *Camion) {
	fmt.Printf("%v GONE %v/%v\n", truck.Get_name(), truck.Get_current_weight(), truck.Get_max_weight())
}

func PrintTransPickup(trans *Transpalette, pack *Colis) {
	packWeight := pack.Get_current_weight()
	fmt.Printf("%v TAKE %v %v\n", trans.Get_name(), pack.Get_name(), colors[packWeight])
}

func PrintTransDrop(trans *Transpalette, pack *Colis) {
	packWeight := pack.Get_current_weight()
	fmt.Printf("%v LEAVE %v %v\n", trans.Get_name(), pack.Get_name(), colors[packWeight])
}

func PrintTransWaiting(trans *Transpalette) {
	fmt.Printf("%v WAIT\n", trans.Get_name())
}

func PrintTruckWaiting(truck *Camion) {
	fmt.Printf("%v WAITING %v/%v\n", truck.Get_name(), truck.Get_current_weight(), truck.Get_max_weight())
}

func PrintError(strErr error) {
	fmt.Println(strErr)
}
