package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("My weight on the surface of %v is %v kg.\n", "Earth", 90)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	num := rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	third := 1.0 / 3
	fmt.Println(third)           //0.3333333333333333
	fmt.Printf("%v\n", third)    //0.333333
	fmt.Printf("%f\n", third)    //0.333333
	fmt.Printf("%.3f\n", third)  //0.333
	fmt.Printf("%4.2f\n", third) //0.33

	year := 2020
	fmt.Printf("Type %T for %v\n", year, year)

	//hexa decimal
	var red, green, blue uint8 = 0x00, 0x8d, 0xd5   //0, 141, 213
	fmt.Printf("%x %x %x\n", red, green, blue)      //0 8d d5
	fmt.Printf("#%02x%02x%02x\n", red, green, blue) //#008dd5

	var r uint8 = 255
	r++
	fmt.Println(r) //0

	var number int8 = 127
	number++
	fmt.Println(number) //-128

	var gr uint8 = 3
	fmt.Printf("%08b\n", gr) //00000011

}
