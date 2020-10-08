package main

import(
	"fmt"
	"time"
)

func main(){
	var i int
	var secs int
	secs=3
	for start := time.Now(); time.Since(start) < time.Second*time.Duration(secs); {
		i++
	}
	fmt.Println("a")

}