package main

import "fmt"

func saludos() {
	fmt.Println("Ejecutando linaea")
}
func main() {
	defer saludos() //cambia de prioridad de ejecucion, es decir lo va ejecutar al ultimo
	fmt.Println("saludadno carnal")
}
