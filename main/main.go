package main

import "fmt"
func main (){
	//fmt.Println("Hello, World!");
	fmt.Print("go"+"lang\n");
	fmt.Println("1+1 =", 1+1, "\n");
	fmt.Println("0.7/0.3=", 0.7/0.3);
	fmt.Println(true && false);

	//variables
	var a = "Yash";
	fmt.Println("Hello,", a);
	var b int = 42 ;
	fmt.Println(b);
	c := 3.14;
	c = 34.56
	fmt.Println(c);

	//Constants
	const pi = 3.14159
	fmt.Println("Pi =", pi);

	//loops 
	// while
	i:= 1;
	for i< 12 {
		fmt.Println(i);
		i++;
	}
	
	// for 
	for a := 0 ; a < 4 ; a++{
		fmt.Println("a =", a);
	} 

	for a := 10 ; a>0 ; a--{
		if(a == 2){
			break;
		}
		if (a == 5) {
			continue;
		}
		fmt.Println(a);
	}

	//for range 
	fmt.Println("For Range:");
	for  r := range 10{
		fmt.Println(r);
	}
	for index , value := range  []string{"Go", "is", "awesome"} {
		fmt.Println("index:", index, "value:", value);
	}


}			