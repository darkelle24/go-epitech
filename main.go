package main

import (
	"fmt"

	"github.com/darkelle24/go-epitech/game"
)

func main() {
	var gameEnv game.Game
	gameEnv.Create_map(10, 10)
	gameEnv.Create_camion("test_camion", 1, 1, 1000, 5)
	gameEnv.Create_camion("test_camion_2", 1, 5, 1000, 5)
	gameEnv.Create_camion("test_camion_3", 5, 1, 1000, 5)
	fmt.Println(gameEnv, gameEnv.Map[1][1].Tool.Get_name())
}
