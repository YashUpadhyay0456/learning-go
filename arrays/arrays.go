package main

import "fmt"

func main() {
	fmt.Println("Hi lets learn arrays in Go !")

	var nums [4]int
	for i := range 4 {
		nums[i] = 4 * i
	}
	for i := range 4 {
		fmt.Println(nums[i])
	}

	fmt.Println("nums entire value: ", nums)

	var bools [5]bool
	for i := range 5 {
		bools[i] = i&1 == 0
	}
	fmt.Println(bools)

	var names [3]string
	names[0] = "Yash"
	names[1] = "jayesh"
	fmt.Println(names)
	var empty [12]string
	fmt.Println(empty)
	delarr := [5]int{1, 2, 3, 4, 5}
	fmt.Print(delarr)

}
