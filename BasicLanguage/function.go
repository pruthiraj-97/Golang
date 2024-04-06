package main

import (
	"fmt"
	"reflect"
)

// not allow function inside a function
func greeter()  {
	fmt.Println("hello golang")
}

func addition(num1 int,num2 int) (int,string,int) {
	return num1+num2,"pruthiraj",5
}
func proAdder(values ...int) int{
   values=append(values, 6,7)
   total:=0
   fmt.Println("type of values is: ",reflect.TypeOf(values)) // it is a slice
   for _,value:=range values{
	total+=value
   }
   return total
}

func main() {
	fmt.Println("function")
	var ans,_,num=addition(1,2)
	fmt.Println(ans,num)
	fmt.Println(proAdder(1,2,3,4,5))

}