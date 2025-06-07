package main

import "fmt"

/*

	formatters:
	- %s -> strings
	- %d -> digitos (int..)
	- %f -> float
	- %t -> booleanos
*/

func SomePractice() {

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

	fmt.Printf("los resultados de la suma y multiplicacion de 3 y 3 son: suma: %d, mult: %d\n", s, p)

	fmt.Println("********* SWITCHES *********")
	//switches
	day := "Miercoles"

	switch day {
	case "Lunes":
		fmt.Println("Comienza la semana... :(")
	case "Martes", "Miercoles", "Jueves": //esta es una de las formas de evaluar multiples casos
		fmt.Println("Sobreviviendo....")
	case "VIernes":
		fmt.Println("It's Friday dayyyy")
	default:
		fmt.Println("Sabado o domingo")

	}

	//FORLOOPs
	fmt.Println("********** FOR LOOPS **********")

	for i := 0; i < 5; i++ {
		fmt.Printf("Iteracion %d\n", i+1)
	}

	//WHILE
	//no existe el while, es un for solo con la condicion
	count := 0
	for count < 5 {
		fmt.Println("El counter del while en", count)
		count++
	}

	/*
		//infinite loop
		for{
			//corre hasta que cumpla alguna condicion y le damos al break
			break
		}
	*/

	//arrays y slices
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("El array es", numbers, "tiene una longitud de ", len(numbers))
	//no se pueden agregar mas elementos que los 5, ni programaticamente
	//no podemos hacer numbers[5] = 6

	//slice
	//puede ser: un array dinamico o una porcion (slice) de un array
	//array dinamico: nos permite agregar elementos y si exceden su capacidad, la duplica
	//para crear un slice a partir de otro slice usamos la sintaxis [inicio:final]

	//convertimos el array numbers en un slice
	numbersSlice := numbers[:]
	//creamos un slice con los primeros 3 numeros
	firstThree := numbers[:3]

	fmt.Println("Numbers slice", numbersSlice)
	fmt.Println("First three", firstThree)
}

func add(a int, b int) int {
	return a + b
}

// si devuelve 2 resultados, especificamos todo en parentesis
func calcSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
