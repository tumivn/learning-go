package main

//heap is a data structure used in selection, graph, and k-way merge algorithms
//operations such as finding, merging, insertion, key changes, and deleting are performed on heap
// the value stored at each node is greater than or equal to its children

type IntegerHeap []int

func (intHeap IntegerHeap) Len() int {return len(intHeap)}

func (intHeap IntegerHeap) Less(i, j int) bool {return intHeap[i] < intHeap[j]}

//Still don't understand the logic behind

func (intHeap IntegerHeap) Swap(i, j int) { intHeap[i], intHeap[j] = intHeap[j], intHeap[i]}

func (intHeap *IntegerHeap) Push(heapInterface interface{}){
	*intHeap = append(*intHeap, heapInterface.(int))
}