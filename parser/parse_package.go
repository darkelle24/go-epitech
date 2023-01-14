package parser

import (
	"fmt"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

func getWeight(color string) (weight int, err error) {
	color = strings.ToLower(color)
	switch color {
	case "blue":
		return 500, nil
	case "green":
		return 200, nil
	case "yellow":
		return 100, nil
	}

	return 0, errColor
}

func parserPackage(input string) (name string, posX int, posY int, weight int, err error) {
	var color string

	if strings.Count(input, " ") != 3 {
		return "", 0, 0, 0, errWrongNumberOfParam
	}

	n, err := fmt.Sscanf(input, "%s %d %d %s", &name, &posX, &posY, &color)
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("%w", err)
	}

	if n != 4 {
		return "", 0, 0, 0, errWrongNumberOfParam
	}

	weight, err = getWeight(color)
	if err != nil {
		return "", 0, 0, 0, err
	}

	if posX < 0 || posY < 0 {
		return "", 0, 0, 0, errNegaValue
	}

	return
}

func createPackage(input string, gameEnv *game.Game) error {
	name, posX, posY, weight, err := parserPackage(input)
	if err != nil {
		return err
	}

	if err := gameEnv.CreateColis(name, posX, posY, weight); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
