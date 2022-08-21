package main

import "fmt"

func main() {
	PerfectNum(28)
}

func PerfectNum(num int) {

	var sum int = 0

	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i

		}

	}
	if sum == num {
		fmt.Println("Perfect Number")

	} else {
		fmt.Println("NOt Perfect number")
	}

}
