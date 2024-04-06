package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("get request")
}

func performGetRequest(){
	const myUrl="http://localhost:4000/get"
	response,err:=http.Get(myUrl)

	if err!=nil {
		panic(err)
	}
	fmt.Println(response.Body)
}