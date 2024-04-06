package main

import "fmt"

func main() {
	fmt.Println()
	// slices IMPORTANT for Backend
//    var fruitList = []string{}
// // method to print type
// fmt.Println("Type of fruiteList is: ",reflect.TypeOf(fruitList))
// fruitList=append(fruitList,"mango","banana")
// fmt.Println(fruitList)
// fmt.Println("Length of list is: ",len(fruitList))

// // slice using make() memory allocation
// highScores:=make([]int,0)
// fmt.Println(highScores)
// highScores=append(highScores, 1,2,3,4) // destracture
// fmt.Println(highScores)

// remove data based on index IMPORTANT
   var course=[]string{"react","js","java","python"}
   fmt.Println(course)
   var index=2
   course=append(course[:2],course[index+1:]... )
   fmt.Println(course)
}