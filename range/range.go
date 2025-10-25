package main

import "fmt"

func main() {
	//use for iterating in go
	num := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 12, 13}, {14, 15, 17}}
	fmt.Println(len(num[3]));
	for i := range len(num) {
		for j := range len(num[i]) {
			fmt.Print(num[i][j]," ");
		}
		fmt.Println()
	}
	
	nums := []int{1,2,3}
	// sum:=0
	fmt.Println(nums)
	for i,nums:=range nums{
		fmt.Println(nums, i)
	}
	// fmt.Print(sum)

	// iterating over an map 
	m := map[string]string {"Fname":"John","Lname":"Doe"};
	for i , mi := range m{
		fmt.Println(i,mi);
	}

	// using range over string 
	for i, c:=range "Yash Upadhyay"{
		fmt.Println(i,c)
	}
}