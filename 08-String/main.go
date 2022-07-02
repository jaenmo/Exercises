// Enter the string STR from the keyboard,
// use the for RANGE statement to traverse each character in the string and print the output.

package main

import "fmt"

// func main(){
// s:= "STR"
// for _, v:= range s{
// 	fmt.Printf("%c \n", v)
// }

// }

func main() {
	var str string
	fmt.Scan(&str)
	for _, v := range str {
		fmt.Println(string(v))
	}
}