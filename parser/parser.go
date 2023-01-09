package parser

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/darkelle24/go-epitech/game"
)

func getPath() (string, error) {
	arg_len := len(os.Args[0:])
	if arg_len != 2 {
		fmt.Println("wrong number of argument")
		return "", errors.New("wrong number of argument")
	}

	return os.Args[1], nil
}

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Conversion des bytes en chaîne de caractères
	str := string(data)
	return str, nil
}

func checkNumber(input string) (int, error) {
	if i, err := strconv.ParseInt(input, 10, 64); err == nil {
		return int(i), nil
	} else {
		return 0, err
	}
}

func firstLineParse(line string) (width int, height int, numberTurnSimulate int, err error) {
	argument := strings.Split(line, " ")

	if len(argument) != 3 {
		return 0, 0, 0, errors.New("wrong number of parameters")
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
		return 0, 0, 0, errors.New("the value can t be negative")
	}

	return width, height, numberTurnSimulate, nil
}

func switchParser(state *int, gameEnv *game.Game, s string) error {
	switch *state {
	case 1:
		if width, height, turn, err := firstLineParse(s); err == nil {
			gameEnv.Create_map(width, height)
			gameEnv.Set_turns(turn)
		} else {
			return err
		}
		*state = *state + 1
	case 2:
		list := strings.Split(s, " ")
		if len(list) == 3 {
			*state = *state + 1
			return switchParser(state, gameEnv, s)
		} else if len(list) != 4 {
			return errors.New("wrong number of parameters")
		}
		if err := createPackage(s, gameEnv); err != nil {
			return err
		}
	case 3:
		list := strings.Split(s, " ")
		if len(list) == 5 {
			*state = *state + 1
			return switchParser(state, gameEnv, s)
		} else if len(list) != 3 {
			return errors.New("wrong number of parameters")
		}
		if err := createPalletTruck(s, gameEnv); err != nil {
			return err
		}
	case 4:
		list := strings.Split(s, " ")
		if len(list) != 5 {
			return errors.New("wrong number of parameters")
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
			fmt.Println(err)
			return err
		}
	}

	if len(gameEnv.Transps) == 0 || len(gameEnv.Packs) == 0 || len(gameEnv.Trucks) == 0 {
		fmt.Println("need min 1 pallet truck, min 1 truck and min 1 package")
		return errors.New("need min 1 pallet truck, min 1 truck and min 1 package")
	}

	return nil
}

func Parser(gameEnv *game.Game) error {

	path, err := getPath()

	if err != nil {
		return err
	}

	file, err := readFile(path)

	if err != nil {
		return err
	}

	// fmt.Println(file)

	fileArray := strings.Split(file, "\n")

	if err := orderParser(fileArray, gameEnv); err != nil {
		return err
	}

	return nil
}
