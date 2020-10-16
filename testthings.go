package main

import(
	"fmt"
)

func main(){
	var stuff []int
	stuff=append(stuff,1)

	a,b:=stuff[0], stuff[1:]
	fmt.Println(a)
	fmt.Println(b)

}