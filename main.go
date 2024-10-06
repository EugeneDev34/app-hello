package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

var usageString = fmt.Sprintf(`Использование: %s <целое число> [-h | --help]
Приложение приветствия, которое выводит введённое вами имя <целое число> раз.`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}
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

func validaterArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("необходимо указать число больше 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil

	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("рад встрече %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func main() {
	c, err := parceArg(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validaterArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Println(os.Stdout, err)
		os.Exit(1)
	}
}
