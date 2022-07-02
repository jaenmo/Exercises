// FIZZBUZZ is a game that is often played in the British school.
// From 1 to 100, if you meet 3 times to say fizz, if you meet 5 times to say buzz,
// if it is both a multiple of 5 times, it is necessary to say FIZZBUZZ .
// Write a program and print a number 1-100 according to the game.

package main

import "fmt"

func main(){

	for i:=1 ; i<=100 ; i++ {
		if i%3 == 0 && i%5 !=0 {
			fmt.Println("fizz", i)
		}
		if i%5 == 0 && i%3 !=0{
			fmt.Println("buzz", i)
		}
		if (i%5 == 0) && (i%3 ==0) {
			fmt.Println("FIZZBUZZ", i)
		}
		if (i%5 != 0) && (i%3 !=0) {
			fmt.Println(i)
		}
	}
}