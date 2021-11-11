package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	var count = 10
	for count>0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(5) == 0 {
			fmt.Println("Failed")
			break
		}
		count--
	}
	if count == 0 {
		fmt.Println("Lift off!")
	}

	for i := 10; i>0; i--{
		fmt.Println(i)
	}

	if num := rand.Intn(3); num == 0{
		fmt.Println("Space Adventures")
	}else if num == 1 {
		fmt.Println("SpaceX")
	}else {
		fmt.Println("Virgin Galactic")
	}
}
