package main

import (
	"fmt"
)
type Number interface{
	int64|float64
}

func genericFunction[K comparable, V float64](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}
	return sum
}

func genericMethod12[K comparable,V Number](m map[K]V) V  {
	var sum V
	for _,value:=range m{
		sum+=value
	}
	return sum
}

func main() {
	fmt.Println("generic function")
	var intMap = make(map[string]float64)
	intMap["one"] = 1
	intMap["two"] = 2
	fmt.Println("Generic Sums: %v\n", genericFunction[string, float64](intMap))
	fmt.Println(genericMethod12(intMap))
}
