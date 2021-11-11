package main

import "fmt"

//tuple is a finite sorted list of elements, immutable sequential collection
func powerSeries(a int)(int, int, error){
	return a*a, a*a*a, nil
}

func main(){
	var square int
	var cube int
	var err error

	square, cube, err = powerSeries(3)

	if err == nil {
		fmt.Println("Square ", square, " Cube ", cube)
	}else {
		fmt.Println("Error ", err)
	}
}