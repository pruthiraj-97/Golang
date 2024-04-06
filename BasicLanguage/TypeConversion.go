package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("")
	fmt.Println("wellcome to my app")
	reader:=bufio.NewReader(os.Stdin)

	ratting,_:=reader.ReadString('\n') // input Type string
	numRatting,err:=strconv.ParseInt(strings.TrimSpace(ratting),10,64)
	if err!=nil {
		fmt.Println(err)
	}else{
	   fmt.Println("ratting is ",ratting)
	   fmt.Println("Adding one to your ratting",numRatting+1)
	}
}