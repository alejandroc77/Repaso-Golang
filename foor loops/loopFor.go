package main

import "fmt"

func main() {
	number := 1
	for i := 0; i < 5; i++ {
		number += i //number = number + 1
		fmt.Println(i)
	}
	fmt.Println("El resultado del bucle es: ", number)

}
