package parser

import (
	"errors"
	"fmt"
	"os"
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

func Parser() error {

	path, err := getPath()

	if err != nil {
		return err
	}

	file, err := readFile(path)

	if err != nil {
		return err
	}

	fmt.Println(file)

	//fileArray := strings.Split(file, "\n")

	return nil
}
