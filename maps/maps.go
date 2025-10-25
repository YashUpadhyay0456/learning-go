package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps is similar to hashmaps in java and cpp
	// creating a map
	m := make(map[string]int)
	// adding elements to map
	m["name"] = 12
	m["nae"] = 13
	fmt.Println(m["name"])
	fmt.Println(m)
	fmt.Println(m["nae"])
	delete(m,"nae")
	clear(m)
	fmt.Println(m);

	// create map without make()
	mp := map[string]int{"name":1,"price":12}
	fmt.Println(mp);

	k,ok := mp["name"]
	fmt.Println(k,ok)
	if (ok){
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

	m1 := map[string]int{"id":1,"code":2};
	m2 := map[string]int{"id":1,"cod":2};
	fmt.Println(maps.Equal(m1,m2));
}