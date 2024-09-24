// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// func imprimir() {
// 	fmt.Println("texto desde la funcion imprimir")
// }

// func hola_string(s string) string {
// 	return "hola" + s
// }

// func sumar(a int, b int) int {
// 	return a + b
// }

// func comparar_bool() {
// 	var a bool
// 	//b := true
// 	a = false
// 	if a == true {
// 		fmt.Println("a y b son iguales")
// 	} else {
// 		fmt.Println("a y b no son iguales")
// 	}
// }

// // func main() {
// //fmt.Println("texto desde la funcion main")
// // imprimir()
// // fmt.Println(hola_string(" felipe"))
// // fmt.Println(sumar(3, 5))
// // comparar_bool()

// func arreglo() {
// 	var aprendices [5]string
// 	aprendices[0] = "felipe"
// 	aprendices[1] = "juan"
// 	aprendices[2] = "frank"
// 	aprendices[3] = "jose"
// 	aprendices[4] = "jaider"

// 	fmt.Println(aprendices[1])

// }

// func main() {
// 	// arreglo()
// 	// tipo_datos()
// 	//convertir_string_to_bool()
// 	// convertir_bool_to_string()
// }

// func tipo_datos() {
// 	var texto string = "felipe"
// 	var numero int = 1000
// 	var decimal float64 = 0.00001
// 	fmt.Println(reflect.TypeOf(texto))
// 	fmt.Println(reflect.TypeOf(numero))
// 	fmt.Println(reflect.TypeOf(decimal))

// }

// func convertir_string_to_bool() {
// 	var palabra string = "true"

// 	boolean, err := strconv.ParseBool(palabra)
// 	fmt.Println(boolean, reflect.TypeOf(boolean))
// 	fmt.Println("este es el error", err)
// }

// func convertir_bool_to_string() {
// 	var palabra_bool bool = true

//		strbool := strconv.FormatBool(palabra_bool)
//		fmt.Println(strbool, reflect.TypeOf(strbool))
//	}
package main

import "fmt"

// //func main() {
// 	var (
// 		nombre     string = "nicolas"
// 		edad       int    = 20
// 		pensionado bool   = true
// 	)
// 	fmt.Println("Nombre: ", nombre)
// 	fmt.Println("Edad: ", edad)
// 	fmt.Println("Pensionado: ", pensionado)
// }

// func main() {
// 	var nombre string
// 	var edad int
// 	var peso float64
// 	var estudiante bool
// 	fmt.Println("Nombre: ", nombre)
// 	fmt.Println("Edad: ", edad)
// 	fmt.Println("Peso: ", peso)
// 	fmt.Println("Estudiante: ", estudiante)
// }
func main() {
	nombre := "Benjamin Button"
	edad := 99
	peso := 80
	estudiante := false
	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)
	fmt.Println("Peso: ", peso)
	fmt.Println("Estudiante: ", estudiante)
}
