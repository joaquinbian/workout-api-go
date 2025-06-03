package main

import "fmt"

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
}
