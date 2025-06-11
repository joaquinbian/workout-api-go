package main

import (
	"fmt"
)

/*
formatters:
- %s -> strings
- %d -> digitos (int..)
- %f -> float
- %t -> booleanos
*/
type Person struct {
	Name string
	Age  int
}

// method receiver
func (p *Person) modifyPersonName(name string) {
	p.Name = name
}

func (p Person) sayName() {
	fmt.Printf("Mi nombre es %s\n", p.Name)
}
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

	//para crear un slice no especificamos el tamanio(capacidad )
	fruits := []string{}

	fruits = append(fruits, "Banana")
	fruits = append(fruits, "Manzana")
	fmt.Println("Frutas", fruits)

	//tambien se puede usar con mas de un valor
	fruits = append(fruits, "Durazno", "Mandarina")

	fmt.Printf("Mas frutas %v\n", fruits)

	//tambien se pueden concatenar 2 slices usando el 'spread operator' -> ...

	frutasFavoritasAgus := []string{"Pomelo", "Naranja", "Frutilla"}
	fruits = append(fruits, frutasFavoritasAgus...)

	fmt.Printf("Mis frutas y las frutas favoritas de agus %v\n", fruits)

	//iterando un slice/array
	for index, fruit := range fruits {
		fmt.Printf("el elemento %d de mis frutas es %s\n", index, fruit)
	}

	//Maps
	//key value store
	//sintaxis map[keyType]valueType
	capitales := map[string]string{
		"USA":       "Washington",
		"Argentina": "Buenos Aires",
		"Italia":    "Roma",
		"Chile":     "Santiago",
	}

	//accediendo a un valor de map
	fmt.Printf("la capital de Argentina es %v\n", capitales["Argentina"])
	fmt.Printf("la capital de Brasil es %v\n", capitales["Brasil"]) //nos da el zero value

	//chequear si existe el valor
	capital, ok := capitales["Brasil"]

	if ok {
		fmt.Printf("La capital es %s\n", capital)
	} else {
		fmt.Printf("La capital no existe\n")

	}
	//iterando sobre un map
	for pais, capital := range capitales {
		fmt.Printf("pais %s capital %s\n", pais, capital)
	}

	//eliminamos un valor del map
	delete(capitales, "Chile")

	fmt.Println("Iteramos las capitales sin Chil")
	for pais, capital := range capitales {
		fmt.Printf("pais %s capital %s\n", pais, capital)
	}

	//STRUCTS
	//data type que puede guardar datos y/para usarlos en la app
	structs()
}

func add(a int, b int) int {
	return a + b
}

// si devuelve 2 resultados, especificamos todo en parentesis
func calcSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func structs() {

	yo := Person{
		Name: "Joaquin",
		Age:  24,
	}

	fmt.Printf("Este soy yo %v\n", yo) //{Joaquin 24}

	//para logear tmb las keys agregamos el + al formatter
	fmt.Printf("Este soy yo %+v\n", yo) //{Name:Joaquin Age:24}

	//anonymous struct
	empleado := struct {
		name string
		id   int
	}{
		name: "Joaco",
		id:   12,
	}

	fmt.Printf("Este soy yo empleado%v\n", empleado)         //{Joaco 12}
	fmt.Printf("Este soy yo empleado + keys%+v\n", empleado) //{name:Joaco id:12}

	//cuando pasamos un struct a una funcion, lo pasamos por valor(pasamos una copia a nuestra funcion)
	//y esa copia vive solo en el ambiente de nuestra funcion, esa memoria se libera luego
	//tambien podemos pasar una referencia a un struct e incluso esa se puede modificar

	type Address struct {
		street string
		city   string
	}

	type Contact struct {
		name    string
		address Address
		phone   string
	}

	//se pueden omitir valores al crear un struct gracias a los zero values
	contact := Contact{
		name: "Juan",
		address: Address{
			street: "Calle falsa 123",
			city:   "Ciudad",
		},
	}

	fmt.Printf("Este es mi contacto %v\n", contact)

	//PUNTEROS Y COPIAS, PASAR POR REFERNCIA VS VALOR

	//con el &accedemos a la direccion de memoria de la variable
	//con el *tipo decimos que vamos a manejar un puntero, que en esa variable
	//guardaremos una dir de memoria
	fmt.Printf("Nombre antes de pasar a la funcion: %s\n", yo.Name)
	modifyPersonName(&yo, "Joaquin Fidel")
	fmt.Printf("Nombre depsues de pasar a la funcion %s\n", yo.Name)

	x := 20
	//var xp *int = &x
	xp := &x

	fmt.Printf("Valor de x: %d y direccion de x: %p\n", x, xp)

	//derefernce
	//accedemos al valor del puntero para cambiarlo
	*xp = 30

	fmt.Printf("Valor de x: %d y direccion de x: %p\n", x, xp)

	//methods receivers, metodos de un struct
	yo.modifyPersonName("JoaquinBian")
	yo.sayName()
}

func modifyPersonName(p *Person, name string) {
	p.Name = name
}
