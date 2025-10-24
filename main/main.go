package main

import (
	"fmt"
	"time"
)
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

	//switch case
	// no need to put break statement
	fmt.Println("Switch Case:");
	data := 2
	switch data {
	case 1:
		fmt.Println("Data is 1");
	case 2:
		fmt.Println("Data is 2");
	case 3:
		fmt.Println("Data is 3");
	default:
		fmt.Println("Data is unknown");
	}

	// there is no ternary operator in go

	// multiple condition switch 
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend");
	default:
		fmt.Println("It's a weekday");	
	}
	fmt.Println(time.Now());

	// type switch
	whoAmI := func (i interface{})  {
		switch t :=i.(type) {
		case int:
			fmt.Println("I'm an int");
		case string:
			fmt.Println("I'm a string");
		default:
			fmt.Printf("Don't know type %T\n", t);
		}
	}

	whoAmI(42);
	whoAmI("hello");
	whoAmI(3.14);	

}			
	