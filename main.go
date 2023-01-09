package main

import (
	"fmt"

	"github.com/darkelle24/go-epitech/algo"
	"github.com/darkelle24/go-epitech/game"
	"github.com/darkelle24/go-epitech/parser"
)

func handlePanics() {
	if err := recover(); err != nil {
		fmt.Println("😱", err)
	}
}

func main() {
	defer handlePanics()
	var gameEnv game.Game

	if parser.Parser(&gameEnv) != nil {
		return
	}

	for !gameEnv.IsDone() {
		for _, trans := range gameEnv.Transps {
			trans.Move(0, 0, &gameEnv.Map)
		}
		for _, truck := range gameEnv.Trucks {
			algo.SendTruck(truck, &gameEnv)
		}
		game.UpdateMap(&gameEnv)
		gameEnv.NextTurn()
	}

	// JUSTIN REMOVE
	// x, y := gameEnv.Transps[0].Get_position()
	// gameEnv.Transps[0].Take(x-1, y, &gameEnv.Map)
	// gameEnv.Trucks[0].Wait()
	// gameEnv.NextTurn()
	// gameEnv.Transps[0].Move(1, 0, &gameEnv.Map)
	// gameEnv.Trucks[0].Wait()
	// gameEnv.NextTurn()
	// gameEnv.Transps[0].Drop(x+2, y, &gameEnv.Map)
	// gameEnv.Trucks[0].Move(0, 0, &gameEnv.Map)
	// gameEnv.NextTurn()

	fmt.Println(gameEnv.EndStateCharacter())

}
