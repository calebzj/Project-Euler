package main

import "fmt"
import "testing"

func main() {
	//fmt.Println(problem1())
	var result = testing.Benchmark(BenchmarkMyFunc)
	fmt.Println(result)
}

//BenchmarkMyFunc is testing
func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		problem1()
	}
}

// Multiples of 3 and 5
func problem1() []int {

	return Multiples(3, 5, 0, 10000)
}

// Multiples of 3 and 5
func problem2() []int {
	var results []int
	//var size int = 100
	// for index := 0; index < 1000; index++ {
	// 	go func ()  {
	// 		Multiples(3, 5, index , index + size)
	// 	}
	// }
	return results
}

//Multiples Checks for numbers that are multiples
//of either multiple1 or multiple2 that are less than max
func Multiples(multiple1 int, multiple2 int, min int, max int) []int {
	var results []int
	for index := min; index < max; index++ {
		if index%multiple1 == 0 || index%multiple2 == 0 {
			results = append(results, index)
		}
	}
	return results
}
