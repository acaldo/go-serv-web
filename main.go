package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
	operador  string
	operador1 int
	operador2 int
}

func (c calc) operate(entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		return (operador1 + operador2)
	case "-":
		return (operador1 - operador2)
	case "*":
		return (operador1 * operador2)
	case "/":
		return (operador1 / operador2)
	default:
		fmt.Println("operador no valido")
		return 0
	}
}

func parsear(entrada string) int {
	operador1, _ := strconv.Atoi(entrada)
	return operador1
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()
	return entrada
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	operacion := calc{operador: operador}
	fmt.Println(operacion.operate(entrada, operador))
}
