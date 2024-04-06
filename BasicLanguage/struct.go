package main

import (
	"fmt"
)
type User struct{
	Name string
	Email string
	Status bool
	Age int
}
func main() {
	fmt.Println("struct in golang")
	// no inharitance in golang
	// No super or parent
	pruthiraj:=User{"pruthi","pruthi@gmail.com",true,20}
	fmt.Println(pruthiraj)
	fmt.Printf("user details %+v\n", pruthiraj)
	fmt.Println(pruthiraj.Name,pruthiraj.Age,pruthiraj.Email)

}
