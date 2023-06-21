package main

import (
	"fmt"
	//"math"
	"math/rand"
)

func main() {
	var numeroJugador int
	var numeroRandom = GenerarNumero(1, 10)
	fmt.Println("vamos a tirar los dados")
	fmt.Println("escoge un numero del 1 al 10")
	fmt.Scan(&numeroJugador)
	fmt.Println(numeroRandom)
	if numeroJugador == numeroRandom {
		println("has ganado")
	} else {
		println("has perdido")
	}

}

func GenerarNumero(minimo int, maximo int) int {

	numeroRandom := rand.Intn(maximo-minimo) + minimo

	return numeroRandom
}
