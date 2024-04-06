package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array declaration
    var array = [5]int64{1,2,3,4,5}
    fmt.Println(array)

	// slices IMPORTANT for Backend
	var fruitList = []string{}
	// method to print type
	fmt.Println("Type of fruiteList is: ",reflect.TypeOf(fruitList))
	fruitList=append(fruitList,"mango","banana")
	fmt.Println(fruitList)
	fmt.Println("Length of list is: ",len(fruitList))
    
	// slice using make() memory allocation
	highScores:=make([]int,0)
	fmt.Println(highScores)
    highScores=append(highScores, 1,2,3,4) // destracture
	fmt.Println(highScores)
}