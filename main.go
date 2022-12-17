package main

import (
	"fmt"

	"github.com/darkelle24/go-epitech/game"
)

func handlePanics() {
	if err := recover(); err != nil {
		fmt.Println("😱", err)
	}
}

func main() {
	defer handlePanics()
	var gameEnv game.Game
	fmt.Println(gameEnv)
	gameEnv.Create_map(10, 10)
	gameEnv.Create_camion("test_camion", 1, 1, 1000, 5)
	gameEnv.Create_colis("test_colis", 5, 2, 200)
	gameEnv.Create_transpallete("test_transpallete", 5, 1)
	// fmt.Println(gameEnv, gameEnv.Map[1][1].Tool.Get_name(), gameEnv.Map[1][5].Tool.Get_name(), gameEnv.Map[5][1].Tool.Get_name())
	s, ok := gameEnv.Map[5][1].Tool.(*game.Transpalette)
	truck := gameEnv.Map[1][1].Tool.(*game.Camion)
	fmt.Println(s, ok, s.Get_name())
	var floor = &gameEnv.Map
	fmt.Println((*floor)[5][1].Tool)
	gameEnv.Next_turn()
	s.Take(5, 2, floor)
	s.Move(-1, 0, floor)
	s.Move(-1, 0, floor)
	s.Move(-1, 0, floor)
	fmt.Println(s, truck)
	s.Drop(1, 1, floor)
	fmt.Println(s, truck)
	gameEnv.Next_turn()
	err := s.Move(-1, 0, floor)
	fmt.Println(err)
	gameEnv.Next_turn()
}
