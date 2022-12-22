package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

func parserTruck(input string) (name string, posX int, posY int, weightMax int, turn int, err error) {
	if strings.Count(input, " ") != 4 {
		return "", 0, 0, 0, 0, errors.New("wrong number of parameters for package")
	}

	n, err := fmt.Sscanf(input, "%s %d %d %d %d", &name, &posX, &posY, &weightMax, &turn)

	if err != nil {
		return
	}

	if n != 5 {
		return "", 0, 0, 0, 0, errors.New("wrong number of parameters for package")
	}

	if posX < 0 || posY < 0 || turn < 0 {
		return "", 0, 0, 0, 0, errors.New("the value can t be negative")
	}

	if weightMax < 100 {
		return "", 0, 0, 0, 0, errors.New("weight max can't be lower than 100")
	}

	return
}

func createTruck(input string, gameEnv *game.Game) error {
	name, posX, posY, weightMax, turn, err := parserTruck(input)

	if err != nil {
		return err
	}

	if err = gameEnv.Create_camion(name, posX, posY, weightMax, turn); err != nil {
		return err
	}

	return nil
}
