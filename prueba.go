package main

import "fmt"

func main() {
	var numeroUno int
	var numeroDos int
	fmt.Println("Introduce el primer número")
	fmt.Scan(&numeroUno)
	fmt.Println("Introduce el segundo número")
	fmt.Scan(&numeroDos)
	fmt.Println(SumarNumeros(numeroUno, numeroDos))
}

func SumarNumeros(n1 int, n2 int) int {
	return n1 + n2
}
