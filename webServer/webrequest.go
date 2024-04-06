package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)
const url="https://loc.dev"

func main() {
	fmt.Println("web request")
	response,err:=http.Get(url) // response -> *http.Response
	HandleError(err)
	fmt.Println(reflect.TypeOf(response))
	// response should be close
	// caller responsibility to close it

	defer response.Body.Close()

	dataByte,err:=ioutil.ReadAll(response.Body)
	HandleError(err)
	content:=string(dataByte)
	fmt.Println(content)
}

func HandleError(err error)  {
	if err!=nil {
		panic(err)
	}
}