package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

func parserPalletTruck(input string) (name string, posX int, posY int, err error) {
	if strings.Count(input, " ") != 2 {
		return "", 0, 0, errors.New("wrong number of parameters for package")
	}

	n, err := fmt.Sscanf(input, "%s %d %d", &name, &posX, &posY)

	if err != nil {
		return
	}

	if n != 3 {
		return "", 0, 0, errors.New("wrong number of parameters for package")
	}

	if posX < 0 || posY < 0 {
		return "", 0, 0, errors.New("the value can t be negative")
	}

	return
}

func createPalletTruck(input string, gameEnv *game.Game) error {
	name, posX, posY, err := parserPalletTruck(input)

	if err != nil {
		return err
	}

	if err = gameEnv.Create_transpallete(name, posX, posY); err != nil {
		return err
	}

	return nil
}
