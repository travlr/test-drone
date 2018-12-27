package main

import "fmt"
import "tmp/test-drone/math"

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	avg := math.Average(xs)
	fmt.Println(avg)
}
