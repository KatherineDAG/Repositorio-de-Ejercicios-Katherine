/*
Ejercicio 1: Haga un programa que cuente e indique el número de caracteres, el número de palabras
y el número de líneas de un texto cualquiera (obtenido de cualquier forma que considere).
Considere hacer el cálculo de las tres variables en el mismo proceso.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var texto string
	var numCaracter int
	var numPalabras int
	var numLineas int

	texto = "Los elefantes tienen una esperanza de vida de\n 60 a 70 años en la naturaleza," +
		"\n aunque algunos han vivido hasta los \n80 años en cautiverio."
	numCaracter = len(texto)                 //Cuenta la longitud de la cadena
	numPalabras = len(strings.Fields(texto)) //Busqueda de palabras en la cadena, concatenar string a int
	numLineas = 1

	for _, linea := range texto { //para recorrer el texto y contar los saltos de linea
		if linea == '\n' {
			numLineas++ //aumenta, por cada '\n' encontrado
		}
	}

	fmt.Println("Texto:", texto)
	fmt.Println("El número de caracteres es: ", numCaracter)
	fmt.Println("El número de palabras es: ", numPalabras)
	fmt.Println("El número de lineas es: ", numLineas)
}
