// Execute the program
package main

import (
	"fmt"

	"github.com/darkelle24/go-epitech/game"
	"github.com/darkelle24/go-epitech/parser"
	"github.com/darkelle24/go-epitech/solver"
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
	orders := solver.SetupManager(&gameEnv)

	for !gameEnv.IsDone() {
		orders = solver.UpdateManager(&gameEnv, orders)
		_ = solver.UpdateTrucks(&gameEnv)
		gameEnv.NextTurn()
	}
	fmt.Println(gameEnv.EndStateCharacter())
}
