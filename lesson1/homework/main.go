package main

import (
	"fmt"
	"lecture01_homework/fizzbuzz"
)

func main() {
	for i := 0; i < 101; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i))
	}
}
