package fizzbuzz

import (
	"fmt"
)

func FizzBuzz(i int) string {
	if i/3 == 1 {
		return "Fizz"
	} else if i/5 == 1 {
		return "Buzz"
	} else if i/15 == 1 {
		return "FizzBuzz"
	} else {
		return fmt.Sprint(i)
	}
}
