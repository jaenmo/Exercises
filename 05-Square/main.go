// If a number is equal to the right end of its square, the number is called the same configuration.
// As 5, the number of 25, 5 is the number of right ends in 25, 5 is the same configuration.
// Find the same configuration between 1-1000.

package main

import "fmt"

func main(){

	for i:=1 ; i<=10000 ; i++{
		n := 10
		if i > 10{
			n = 100
		}
		if i > 100{
			n = 1000
		}
		if i*i%n == i {
			fmt.Println(i)
		}
	}
		fmt.Println("25*25=", 25*25)
		fmt.Println("76*76=", 76*76)
		fmt.Println("376*376=", 376*376)
		fmt.Println("625*625=", 625*625)
}