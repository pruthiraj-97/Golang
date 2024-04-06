// Handaling URL in golang
package main

import (
	"fmt"
	"net/url"
	"reflect"
)

const MyUrl = "https://loc.dev:3000/learn?coursename=react&language=golang"

func main() {
	fmt.Println(MyUrl)

	// parsing
	result, _ := url.Parse(MyUrl)
	fmt.Println(result.Scheme) // https
	fmt.Println(result.Path)   // api path /learn
	fmt.Println(result.Port()) // 3000
	fmt.Println(result.RawQuery)

	qparams := result.Query() // key value of params
	for key, value := range qparams {
		fmt.Println(key, value)
	}
	// type is slice
	fmt.Println("Type of value is ", reflect.TypeOf(qparams["coursename"]))

	// creating url
	// parsUrl type is *url.URL
	parsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=pruthi",
	}
	anotherUrl := parsOfUrl.String()
	fmt.Println(reflect.TypeOf(anotherUrl))

}
