package main

import "fmt"

func futbol(player *string) {
	*player = "Messi"
	fmt.Println("El mejor jugador es", *player)
}
func portero(nombre *string) {
	*nombre = "Ter Stegen"
	//fmt.Println("El mejor portero es ", nombre)
	//fmt.Println(nombre)
}
func main() {
	jugador2 := "Cristiano"
	fmt.Println(jugador2)
	//futbol(&jugador2)
	portero(&jugador2)
	//jugador4 := "El bucho"
	fmt.Println(jugador2)
}
