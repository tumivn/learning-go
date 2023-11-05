package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "0b111100100"
	int1, int1err := strconv.ParseInt(val1, 0, 8)
	int2, int2err := strconv.ParseInt(val1, 0, 64)
	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Int1 Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
	if int2err == nil {
		fmt.Println("Int2 Parsed value:", int2)
	} else {
		fmt.Println("Cannot parse", val1, int2err)
	}
}
