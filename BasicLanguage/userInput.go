package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	wellcome:="well come to userInput"
	fmt.Println(wellcome)
	// creating reader
	reader:=bufio.NewReader(os.Stdin)
    fmt.Println("enter ratting")

	//input err
	ratting,_:=reader.ReadString('\n')

	fmt.Println("ratting is",ratting)
}