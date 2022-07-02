// Enter a three-digit integer to find the sum of the numbers on each bit. Such as 123,
// the number of numbers on each bit is 1 + 2 + 3 = 6.

package main

import "fmt"

func main(){
	
fmt.Println("precio de la gasolina en €/l")
var petrol float64
fmt.Scanln(&petrol)

fmt.Println("Días de trabajo")
var day int
fmt.Scanln(&day)
km := 82
t := km * day
fmt.Printf("Km recorridos = %v\n", t)

//consumicion 7 l/100

p := petrol * 7 / 100
p1 := p * float64(t)
fmt.Printf("Gastos de gasolina = %v", p1)

}