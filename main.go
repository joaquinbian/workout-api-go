package main

import "fmt"

/*

	formatters:
	- %s -> strings
	- %d -> digitos (int..)
	- %f -> float
	- %t -> booleanos
*/

func main() {

	var name string = "Joaquin"
	fmt.Printf("hello %s\n", name)

	age := 24
	fmt.Printf("my name is %d\n", age)

	var city string
	city = "Buenos Aires"
	fmt.Printf("hello %s\n", city)

	//multiples variables
	var country, continent = "Buenos Aires", "South America"
	fmt.Printf("%s, %s\n", country, continent)

	//multiples variables de diferentes tipos
	var (
		isEmployed bool   = false
		salary     int    = 100
		role       string = "Developer"
		message    string = ""
	)

	if isEmployed {
		message = fmt.Sprintf("Su trabajo es de %s y gana %d", role, salary)
	} else {
		message = fmt.Sprintf("Su trabajo era %s y ganaba %d", role, salary)

	}

	fmt.Println(message)

	//zero values (valores por defecto)
	//previenen nil errors
	var defaultInt int       //0
	var defaultFloat float64 //0.0
	var defaultString string //""
	var defaultBool bool     //false

	fmt.Printf("Default: %d, %f, %s, %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	//constantes
	const pi = 3.14

	const (
		Lunes     = 1
		Martes    = 2
		Miercoles = 3
	)

	const typedPi float64 = 3.14

	const (
		Enero = iota //0
		Feb          //1
		Mar          //2
		Abr          //3
		May          //4
	)

	fmt.Printf("Meses: %d, %d, %d, %d, %d\n", Enero, Feb, Mar, Abr, May)
	add(2, 3)
	s, p := calcSumAndProduct(3, 3)

	fmt.Printf("los resultados de la suma y multiplicacion de 3 y 3 son: suma: %d, mult: %d", s, p)
}

func add(a int, b int) int {
	return a + b
}

// si devuelve 2 resultados, especificamos todo en parentesis
func calcSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
