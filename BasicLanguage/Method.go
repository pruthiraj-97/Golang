package main

import "fmt"
type User struct{
	Name string
	age int
	isLoggedIn bool
}
// the user that we pass in the copy not the reference 
// use pointer
func (u User) UserDetails(){
	u.Name="Raghav"
    fmt.Println(u)
}

func main() {
	fmt.Println("methos id golang")
    pruthi:=User{
		Name:"pruthiraj",
        age:20,
		isLoggedIn: true,
	   }

	pruthi.UserDetails()
}