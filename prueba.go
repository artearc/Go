package main

import (
	"fmt"
)

// func main() {
// 	// var numeroUno float64
// 	// var numeroDos float64
// 	// fmt.Println("Introduce el primer número")
// 	// fmt.Scan(&numeroUno)
// 	// fmt.Println("Introduce el segundo número")
// 	// fmt.Scan(&numeroDos)
// 	// fmt.Println(SumarNumeros(numeroUno, numeroDos))
// 	// fmt.Println("Introduce el primer número")
// 	// fmt.Scan(&numeroUno)
// 	// fmt.Println("Introduce el segundo número")
// 	// fmt.Scan(&numeroDos)
// 	// fmt.Println(RestarNumerosNaturales(numeroUno, numeroDos))
// 	fmt.Println("My favorite number is", rand.Intn(10))
// 	fmt.Println(math.Pi)
// 	fmt.Println(split(17))
// 	var i, j int = 1, 2
// 	//Dentro de una función, la declaración de asignación corta := se puede usar en lugar de una declaración var con tipo implícito.
// 	//Fuera de una función, cada declaración comienza con una palabra clave (var, func, etc.) y, por lo tanto,
// 	// la construcción := no está disponible.
// 	k := 3
// 	c, python, java := true, false, "no!"

// 	fmt.Println(i, j, k, c, python, java)

// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 	}
// }

func MultiplyNumbers(n1 float64, n2 float64) float64 {
	return n1 * n2
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
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
