package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

func main() {

}
