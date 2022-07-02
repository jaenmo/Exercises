// A ball falls from 100M highly free, and each time it falls back to the original height half,
// then fall, ask how many meters in the 10th floor? How high is the 10th rebound?


package main
 
import "fmt"
 
func main() {
	h := 100.0
	s := 0.0
	for i := 0; i < 10; i++ {
		s += h
		h /= 2.0
		s += h
	}
	s -= h
	fmt.Println(s, h)
}
