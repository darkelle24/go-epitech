// Execute the program
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
	orders := algo.SetupManager(&gameEnv)

	for !gameEnv.IsDone() {
		orders = algo.UpdateManager(&gameEnv, orders)
		_ = algo.UpdateTrucks(&gameEnv)
		gameEnv.NextTurn()
	}
	fmt.Println(gameEnv.EndStateCharacter())
}
