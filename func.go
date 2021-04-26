package main

import (
	"fmt"
    "time"
    "math/rand"
)

func foo (params... int){
	for _, param := range params{
		fmt.Println(param)
	}
}



func main() {
	// foo(1,2,3,5)

	v1 := 11
	if(v1 ==11){
		// fmt.Println("合ってる")
	}

	rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(7)).
	v2:= rand.Intn(7)
	fmt.Println(v2)
}

