package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

func parserTruck(input string) (name string, posX int, posY int, weightMax int, turn int, err error) {
	if strings.Count(input, " ") != 4 {
		return "", 0, 0, 0, 0, errWrongNumberOfParam
	}

	n, err := fmt.Sscanf(input, "%s %d %d %d %d", &name, &posX, &posY, &weightMax, &turn)
	if err != nil {
		return "", 0, 0, 0, 0, errors.Unwrap(err)
	}

	if n != 5 {
		return "", 0, 0, 0, 0, errWrongNumberOfParam
	}

	if posX < 0 || posY < 0 || turn < 0 {
		return "", 0, 0, 0, 0, errNegaValue
	}

	if weightMax < 100 {
		return "", 0, 0, 0, 0, errWeight
	}

	return
}

func createTruck(input string, gameEnv *game.Game) error {
	name, posX, posY, weightMax, turn, err := parserTruck(input)
	if err != nil {
		return err
	}

	if err := gameEnv.CreateCamion(name, posX, posY, weightMax, turn); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}
