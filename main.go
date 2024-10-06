package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Как вас зовут? Когда закончите, нажмите клавишу Enter.\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("Ты не ввел свое имя")
	}
	return name, nil
}

type config struct {
	numTimes   int
	printUsage bool
}

func parceArg(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("неверное количество аргументов ")
	}

	if args[0] == "-h" || args[0] == "- -help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes
	return c, nil
}

func main() {

}
