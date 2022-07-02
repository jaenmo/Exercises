// There are 10 stations on an iron route. If a ticket is needed between the two stations,
// how many tickets are prepared?

package main

import "fmt"

// func main(){
// s:= 10
// t:= 0

// for i:=0; i<=10; i++{
// 		t= s-1
// 	}
// 	fmt.Println(t, s)
// }
 
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			sum++
		}
	}
	fmt.Println(sum)
}