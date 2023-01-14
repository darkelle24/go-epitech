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

	if err := parser.Parser(&gameEnv); err != nil {
		fmt.Println("😱", err)
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
