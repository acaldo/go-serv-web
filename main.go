package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
	operator  string
	operador1 int
	operador2 int
}

func (c calc) operate(entry string, operator string) int {
	cleanEntrance := strings.Split(entry, operator)
	operador1 := parsear(cleanEntrance[0])
	operador2 := parsear(cleanEntrance[1])
	switch operator {
	case "+":
		return (operador1 + operador2)
	case "-":
		return (operador1 - operador2)
	case "*":
		return (operador1 * operador2)
	case "/":
		return (operador1 / operador2)
	default:
		fmt.Println("invalid operator")
		return 0
	}
}

func parsear(entry string) int {
	operador1, _ := strconv.Atoi(entry)
	return operador1
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entry := scanner.Text()
	return entry
}

func main() {
	entry := readEntry()
	operator := readEntry()
	operation := calc{}
	fmt.Println(operation.operate(entry, operator))
}
