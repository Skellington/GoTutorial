package main

import (
	"fmt"
	"math"
)

var a, b, c, d, raizd, x1, x2 float64

func main() {
	fmt.Printf("Olá! Seja bem vindo a calculadora de equações de segundo grau.\n")
	fmt.Printf("Insira o valor de A:\n")
	fmt.Scan(&a)
	fmt.Printf("Insira o valor de B:\n")
	fmt.Scan(&b)
	fmt.Printf("Insira o valor de C:\n")
	fmt.Scan(&c)
	if a == 0 {
		fmt.Printf("Se o coeficiente A for igual a zero, a parábola não pode ser calculada.\n")
		fmt.Printf("Por favor, reinsira os valores dos coeficientes.\n")
		return
	} else {

		d := math.Pow(b, 2) - 4*a*c
		raizd := math.Sqrt(d)
		fmt.Printf("Delta = %v\n", d)

		switch {
		case d < 0:
			fmt.Printf("Outch! :-(\n")
			fmt.Printf("Infelizmente o valor de Delta é negativo. Por favor, reinsira os dados.\n")
		case d == 0:
			x1 = -b / (2 * a)
			fmt.Printf("O valor de X é: %v\n", x1)
		default:
			x1 = (-b + raizd) / (2 * a)
			x2 = (-b - raizd) / (2 * a)
			fmt.Printf("O valor de X' é: %v\n", x1)
			fmt.Printf("O valor de X\" é: %v\n", x2)
		}

	}
}
