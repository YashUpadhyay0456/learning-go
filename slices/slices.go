package main

import (
	"fmt"
	"slices"
)

func main() {
	// nothing but dynamic arrays
	// not requred to give length
	// use when the size is unknown
	// most use construct in go
	// contains usefull methos in go for using slices efficiently
	var nums []int //uninitialsed is nil
	fmt.Println(nums)
	fmt.Println(nums == nil)
	fmt.Println(len(nums))
	var sl = make([]int, 0, 5)
	fmt.Println(sl)
	sl = append(sl, 12)
	sl = append(sl, 13)
	sl = append(sl, 14)
	sl = append(sl, 12)
	sl = append(sl, 13)
	sl = append(sl, 14)
	sl = append(sl, 1)
	fmt.Println(cap(sl))
	fmt.Println(sl)

	fmt.Println(sl == nil)
	

/*	num := []int{}
	num = append(num, 1);
	num = append(num, 2);
	num = append(num, 3);
	// fmt.Printf("cap: %d, len: %d slice num: %v", cap(num), len(num), num);

	// copy funtion 
var numcopy = make([]int,len(num));
	copy(numcopy,num);
	fmt.Println(len(num));
	fmt.Printf("num: %v, numcopy: %v",num,numcopy);
*/
	// slice operator
	var num = []int{1,2,3,4,5,6,7,8,9,0};
	fmt.Println(num[3:]);
	var num2 = []int{1,2,3,4};
	fmt.Println(slices.Equal(num,num2));
// 2d slices 
	var num3 = [][]int{{1,2,3},{4,5,6},{7,8,9}};
	for i:= range 3{
		for j :=range 3{
			fmt.Print(num3[i][j])

		}
		fmt.Println();
	}
}
