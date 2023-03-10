// Package parser implements utility for paser file and check error
package parser

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

// Errors

var (
	errWrongNumberOfArg   = errors.New("wrong number of argument")
	errWrongNumberOfParam = errors.New("wrong number of parameters")
	errNegaValue          = errors.New("the value can t be negative")
	errWrongNumberTurn    = errors.New("wrong number of turn")
	errFile               = errors.New("need min 1 pallet truck, min 1 truck and min 1 package")
	errWeight             = errors.New("weight max can't be lower than 100")
	errColor              = errors.New("problem with color")
)

func getPath() (string, error) {
	argLen := len(os.Args[0:])
	if argLen != 2 {
		return "", errWrongNumberOfArg
	}

	return os.Args[1], nil
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", errors.Unwrap(err)
	}

	str := string(data)
	return str, nil
}

func checkNumber(input string) (int, error) {
	i, err := strconv.ParseInt(input, 10, 64)

	if err == nil {
		return int(i), nil
	}

	return 0, errors.Unwrap(err)
}

func firstLineParse(line string) (width int, height int, numberTurnSimulate int, err error) {
	argument := strings.Split(line, " ")

	if len(argument) != 3 {
		return 0, 0, 0, errWrongNumberOfParam
	}

	width, err = checkNumber(argument[0])

	if err != nil {
		return 0, 0, 0, err
	}

	height, err = checkNumber(argument[1])

	if err != nil {
		return 0, 0, 0, err
	}

	numberTurnSimulate, err = checkNumber(argument[2])

	if err != nil {
		return 0, 0, 0, err
	}

	if numberTurnSimulate < 0 || height < 0 || width < 0 {
		return 0, 0, 0, errNegaValue
	}

	if numberTurnSimulate < 10 || 100000 < numberTurnSimulate {
		return 0, 0, 0, errWrongNumberTurn
	}

	return width, height, numberTurnSimulate, nil
}

func switchParser(state *int, gameEnv *game.Game, s string) error {
	switch *state {
	case 1:
		if width, height, turn, err := firstLineParse(s); err == nil {
			if mapErr := gameEnv.CreateMap(width, height); mapErr != nil {
				return err
			}
			gameEnv.SetTurns(turn)
		} else {
			return err
		}
		*state++
	case 2:
		list := strings.Split(s, " ")
		if len(list) == 3 {
			*state++
			return switchParser(state, gameEnv, s)
		} else if len(list) != 4 {
			return errWrongNumberOfParam
		}
		if err := createPackage(s, gameEnv); err != nil {
			return err
		}
	case 3:
		list := strings.Split(s, " ")
		if len(list) == 5 {
			*state++
			return switchParser(state, gameEnv, s)
		} else if len(list) != 3 {
			return errWrongNumberOfParam
		}
		if err := createPalletTruck(s, gameEnv); err != nil {
			return err
		}
	case 4:
		list := strings.Split(s, " ")
		if len(list) != 5 {
			return errWrongNumberOfParam
		}
		if err := createTruck(s, gameEnv); err != nil {
			return err
		}
	}
	return nil
}

func orderParser(fileArray []string, gameEnv *game.Game) error {
	state := 1

	for _, s := range fileArray {
		if err := switchParser(&state, gameEnv, s); err != nil {
			return err
		}
	}

	if len(gameEnv.Transps) == 0 || len(gameEnv.Packs) == 0 || len(gameEnv.Trucks) == 0 {
		return errFile
	}

	return nil
}

// Parser parse file and check for errors
func Parser(gameEnv *game.Game) {
	path, err := getPath()
	if err != nil {
		panic(err)
	}

	file, err := readFile(path)
	if err != nil {
		panic(err)
	}

	fileArray := strings.Split(strings.ReplaceAll(file, "\r\n", "\n"), "\n")

	if err := orderParser(fileArray, gameEnv); err != nil {
		panic(err)
	}
}
