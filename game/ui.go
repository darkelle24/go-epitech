package game

import (
	"fmt"
)

var colors = map[int]string{
	100: "YELLOW",
	200: "GREEN",
	500: "BLUE",
}

// UpdateMap prints the map and the tools
func UpdateMap(game *Game) {
	game.PrintMap()
	game.PrintTools()
}

// PrintNextTurn prints current game's turn
func PrintNextTurn(turn int) {
	fmt.Println("tour", turn)
}

// PrintTransMove prints the Transpalette's move
func PrintTransMove(trans *Transpalette, x, y int) {
	fmt.Printf("%v GO [%v,%v]\n", trans.GetName(), x, y)
}

// PrintTruckDepart prints that the truck is gone
func PrintTruckDepart(truck *Camion) {
	fmt.Printf("%v GONE %v/%v\n", truck.GetName(), truck.GetCurrentWeight(), truck.GetMaxWeight())
}

// PrintTruckGone prints that the truck is still gone
func PrintTruckGone(truck *Camion) {
	fmt.Printf("%v GONE %v/%v\n", truck.GetName(), truck.GetCurrentWeight(), truck.GetMaxWeight())
}

// PrintTransPickup prints that the Transpalette picked up a Colis
func PrintTransPickup(trans *Transpalette, pack *Colis) {
	packWeight := pack.GetCurrentWeight()
	fmt.Printf("%v TAKE %v %v\n", trans.GetName(), pack.GetName(), colors[packWeight])
}

// PrintTransDrop prints that the Transpalette droped a Colis
func PrintTransDrop(trans *Transpalette, pack *Colis) {
	packWeight := pack.GetCurrentWeight()
	fmt.Printf("%v LEAVE %v %v\n", trans.GetName(), pack.GetName(), colors[packWeight])
}

// PrintTransWaiting prints that the Transpalette is waiting
func PrintTransWaiting(trans *Transpalette) {
	fmt.Printf("%v WAIT\n", trans.GetName())
}

// PrintTruckWaiting prints that the Camion is waiting
func PrintTruckWaiting(truck *Camion) {
	fmt.Printf("%v WAITING %v/%v\n", truck.GetName(), truck.GetCurrentWeight(), truck.GetMaxWeight())
}

// PrintError prints an error
func PrintError(strErr error) {
	fmt.Println(strErr)
}
