package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

func parserPalletTruck(input string) (name string, posX int, posY int, err error) {
	if strings.Count(input, " ") != 2 {
		return "", 0, 0, errWrongNumberOfParam
	}

	n, err := fmt.Sscanf(input, "%s %d %d", &name, &posX, &posY)
	if err != nil {
		return
	}

	if n != 3 {
		return "", 0, 0, errWrongNumberOfParam
	}

	if posX < 0 || posY < 0 {
		return "", 0, 0, errNegaValue
	}

	return
}

func createPalletTruck(input string, gameEnv *game.Game) error {
	name, posX, posY, err := parserPalletTruck(input)
	if err != nil {
		return err
	}

	if err := gameEnv.CreateTranspallete(name, posX, posY); err != nil {
		return errors.Unwrap(err)
	}

	return nil
}
