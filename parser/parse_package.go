package parser

import (
	"errors"
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

	return 0, errors.New("problem with color")
}

func parserPackage(input string) (name string, posX int, posY int, weight int, err error) {
	var color string

	if strings.Count(input, " ") != 3 {
		return "", 0, 0, 0, errors.New("wrong number of parameters for package")
	}

	n, err := fmt.Sscanf(input, "%s %d %d %s", &name, &posX, &posY, &color)

	if err != nil {
		return "", 0, 0, 0, err
	}

	if n != 4 {
		return "", 0, 0, 0, errors.New("wrong number of parameters for package")
	}

	weight, err = getWeight(color)
	if err != nil {
		return "", 0, 0, 0, err
	}

	if posX < 0 || posY < 0 {
		return "", 0, 0, 0, errors.New("the value can t be negative")
	}

	return
}

func createPackage(input string, gameEnv *game.Game) error {
	name, posX, posY, weight, err := parserPackage(input)

	if err != nil {
		return err
	}

	if err = gameEnv.Create_colis(name, posX, posY, weight); err != nil {
		return err
	}

	return nil
}
