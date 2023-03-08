/*
Ejercicio 5: Amplie el funcionamiento del ejercicio de Productos visto en clase para que el
programa ahora permita:
	a-A partir de la lista de productos con mínimas existencias de inventario, ampliar el inventario
	con la compra de más unidades de cada producto de esta lista hasta que cumplan con el mínimo establecido.
	Se sugiere crear una función denominada “func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos)”
	b-Crear una función que ordene la lista de productos usando como llave para el ordenamiento cualquiera de
	los elementos de la estructura producto. La lista/slice debe quedar modificada al finalizar el método.
	Se solicita investigar y hacer úso de alguna función predefinida de alguna librería del lenguaje Go que
	ayude a resolver el problema.
*/

package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var listaProductosMinimos listaProductos
var lProductos listaProductos

const existenciaMinima int = 10 //La existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

// Agrega productos a la lista
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	productoBuscado := l.buscarProducto(nombre)
	if productoBuscado != -1 { // cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad de elementos del producto y eventualmente el precio si es que es diferente.
		(*l)[productoBuscado].cantidad += cantidad
		(*l)[productoBuscado].precio = precio
		println("El producto ya existe por lo que ha sido modificada su cantidad y precio")
		fmt.Println(lProductos)
	} else {
		lProductos = append(lProductos, producto{nombre: nombre, cantidad: cantidad, precio: precio})
		fmt.Println(lProductos)
	}
}

// Busca un producto por nombre.
func (l *listaProductos) buscarProducto(nombre string) int { // El retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i // Se iguala a la posicion en la lista donde se encuentra el producto
		}
	}
	return result
}

// Elimina los datos que tengan de cantidad cero en la lista
func (lista *listaProductos) eliminarProducto() {
	var i int
	for i = 0; i < len(*lista); i++ {
		if (*lista)[i].cantidad == 0 {
			*lista = append((*lista)[:i], (*lista)[i+1:]...)
			println("Producto Eliminado")
		}
	}
	fmt.Println(lProductos)
}

// Vende los productos siempre y cuando haya la cantidad correcta.
func (l *listaProductos) venderProducto(nombre string, cant int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cant > 0 {
		if (*l)[prod].cantidad >= cant {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cant
			fmt.Println("Producto vendido: ", (*l)[prod].nombre)
		} else {
			l.eliminarProducto()
			fmt.Println("No se puede vender mayor cantidad de productos de que los que hay en existencia")
		}
	}
}

// Agregar los productos
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("galletas", 50, 1750)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	var index int
	for index = 0; index < len(*l); index++ {
		if (*l)[index].cantidad < existenciaMinima {
			listaProductosMinimos = append(listaProductosMinimos, (*l)[index])
		}
	}
	fmt.Println("Lista de producto con cantidad por debajo del minimo: ", listaProductosMinimos)
	return listaProductosMinimos
}

// Ordenar la lista por medio de la cantidad existente de cada producto (De menor a mayor)
func (lista *listaProductos) ordenarLista() {
	sort.Slice(lProductos, func(i, j int) bool {
		return lProductos[i].cantidad < lProductos[j].cantidad
	})
	fmt.Println("Lista ordenada del producto con menos cantidad al de mayor cantidad: ", lProductos)
}

// Aumenta las cantidades de los articulos de la lista "listaProductosMinimos"
func (l *listaProductos) aumentarInventarioDeMinimos(nombre string, cantidad int) {
	product := l.buscarProducto(nombre)
	if product != -1 {
		var i int
		for i = 0; i < len(listaProductosMinimos); i++ {
			(*l)[i].cantidad += cantidad
			fmt.Println("El producto:", (*l)[i].nombre, "se le agrego:", cantidad)
		}
	}
	fmt.Println(lProductos)
	l.ordenarLista()
}

func main() {
	llenarDatos()
	lProductos.venderProducto("arroz", 5)
	lProductos.venderProducto("frijoles", 4)
	lProductos.venderProducto("leche", 10)
	lProductos.venderProducto("leche", 50)
	lProductos.ordenarLista()
	lProductos.listarProductosMínimos()
	lProductos.aumentarInventarioDeMinimos("leche", 10)
}
