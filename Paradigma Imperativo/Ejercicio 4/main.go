/*
Ejercicio 4: Probar su funcionamiento creando una lista/slice de personas y filtrando
aquellas personas que sean mayores de edad
*/
package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type persona struct {
	nombre string
	edad   int
}
type Persona []persona

var listaPersona Persona

func (nuevList *Persona) agregarPersona(nombre string, edad int) {
	index := slices.IndexFunc(*nuevList, func(p persona) bool { return p.nombre == nombre })
	if index == -1 {
		*nuevList = append(*nuevList, persona{nombre, edad})
	} else {
		fmt.Println("La persona ya existe en la lista")
	}
}

// Map generico, retorna lista de personas mayores de 18
func map1(lista []persona, k func(per persona) persona) []persona {
	lisFinal := make([]persona, 0, len(lista))

	for _, p := range lista {
		nuevoP := k(p)
		if nuevoP.edad >= 18 {
			lisFinal = append(lisFinal, nuevoP)
		}
	}

	return lisFinal
}

// filter generico, retorna lista de personas mayores de 18
func filter1(persons []persona, filterFunc func(person persona) bool) []persona {
	listaFinal := make([]persona, 0, len(persons))

	for _, p := range persons {
		if p.edad >= 18 {
			listaFinal = append(listaFinal, p)
		}
	}

	return listaFinal
}

// filter generico, retorna lista de personas mayores de 18 pero con filtro de nombre con longitud impar
func filter2(listper []persona, k func(pers persona) bool) []persona {
	listaFinal := make([]persona, 0, len(listper))

	for _, p := range listper {
		if k(p) && p.edad >= 18 {
			listaFinal = append(listaFinal, p)
		}
	}

	return listaFinal
}

func main() {
	listaPersona.agregarPersona("Jared", 12)
	listaPersona.agregarPersona("Anleth", 18)
	listaPersona.agregarPersona("Allison", 21)
	listaPersona.agregarPersona("Romey", 29)

	//	Prueba para el map
	perMayores := map1(listaPersona, func(p persona) persona {
		return p
	})

	//	Prueba para el filter que solo da personas mayores de 18 años
	perMay := filter1(listaPersona, func(p persona) bool {
		return p.edad > 0
	})

	//	Prueba para el filter que solo da personas mayores de 18 años, pero con el filtro de longitud impar de nombres
	perMayNom := filter2(listaPersona, func(p persona) bool {
		return len(p.nombre)%2 == 1
	})

	fmt.Println("Lista Original:")
	fmt.Println(listaPersona)

	fmt.Println("Personas mayores de edad, utilizando el map:")
	fmt.Println(perMayores)

	fmt.Println("Personas mayores de edad, utilizando el filter1:")
	fmt.Println(perMay)

	fmt.Println("Personas mayores de edad, utilizando el filter2:")
	fmt.Println(perMayNom)

}
