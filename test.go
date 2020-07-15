package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	operacion := scanner.Text()
	operador := "-"
	//fmt.Println(operacion)

	switch operador {
	case "+":

	}

	valores := strings.Split(operacion, operador)

	fmt.Println(valores[0], valores[1])

	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])

	fmt.Println(operador1 + operador2)
}
