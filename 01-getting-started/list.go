package main

import (
	list2 "container/list"
	"fmt"
)

func main(){
	var list list2.List
	list.PushBack(11)
	list.PushBack(23)
	list.PushBack(34)

	for element := list.Front(); element != nil; element = element.Next(){
		fmt.Println(element.Value.(int))
	}
}