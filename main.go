package main

import "fmt"

func prueba(a *int, b *string) {
	*a = 645
	*b = *b + "jory"
}

func main() {
	age := 32
	name := "alejandro \t"
	fmt.Printf("Hola que ta mi edad es %v y mi nombre es %v", age, name)

	prueba(&age, &name)
	fmt.Println("\n Salida nueva de ", age, name)

	const numero = 32
	fmt.Println(numero)
}
