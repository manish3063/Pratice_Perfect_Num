package main

import "fmt"

func main() {
	final := PerfectNum(28)

	if final == true {
		fmt.Println("Perfect number")

	} else {
		fmt.Println("not a perfect numb")
	}

}

func PerfectNum(num int) bool {

	result := false
	var sum int = 0

	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i

		}

	}
	if sum == num {
		//fmt.Println("Perfect Number")
		result = true
		return result

	} else {
		//fmt.Println("NOt Perfect number")
		result = false
		return result
	}
	return result

}
