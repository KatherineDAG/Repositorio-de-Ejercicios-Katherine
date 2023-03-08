/*
Ejercicio 2: Escriba el programa más eficiente que pueda para imprimir en pantalla la siguiente figura:
  *
 ***
*****
 ***
  *
Para dicho fin, escriba y use una función que reciba de parámetro la cantidad de elementos de la línea
del centro, la cual debe ser impar positiva.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	//Ingreso del número de asteriscos que desea en la línea del centro de la figura
	var cantElementos int
	fmt.Print("Ingrese la cantidad de elementos que desea para la línea del centro: ")
	fmt.Scan(&cantElementos)

	//Verifica que el número sea impar positivo
	if cantElementos%2 == 0 || cantElementos < 0 {
		fmt.Println("El número ingresado es par o negativo. Solo se aceptan impares positivos.")
		return
	} else {
		recibirParametro(cantElementos)
	}
}

// Recibe número de asteriscos e imprime según iteraciones por medio de operadores
// El strings.Repeat permite que por cada elemento de i iterado se retorna * con cada repeticón indicada
func recibirParametro(cant int) {
	for i := 1; i <= cant; i += 2 {
		fmt.Println(strings.Repeat(" ", (cant-i)/2) + strings.Repeat("*", i) + strings.Repeat(" ", (cant-i)/2))
	}
	for i := cant - 2; i >= 1; i -= 2 {
		fmt.Println(strings.Repeat(" ", (cant-i)/2) + strings.Repeat("*", i) + strings.Repeat(" ", (cant-i)/2))
	}
}
