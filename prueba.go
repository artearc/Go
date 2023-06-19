package main

import "fmt"

func main() {
	var numeroUno float64
	var numeroDos float64
	fmt.Println("Introduce el primer número")
	fmt.Scan(&numeroUno)
	fmt.Println("Introduce el segundo número")
	fmt.Scan(&numeroDos)
	fmt.Println(SumarNumeros(numeroUno, numeroDos))
	fmt.Println("Introduce el primer número")
	fmt.Scan(&numeroUno)
	fmt.Println("Introduce el segundo número")
	fmt.Scan(&numeroDos)
	fmt.Println(RestarNumerosNaturales(numeroUno, numeroDos))
}

func SumarNumeros(n1 float64, n2 float64) float64 {
	return n1 + n2
}
func RestarNumerosNaturales(n1 float64, n2 float64) float64 {
	var resultado float64
	if n1 < n2 {
		fmt.Println("el resultado será negativo")
	} else {
		resultado = n1 - n2
	}
	return resultado
}
