package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime:=time.Now()
	fmt.Println(presentTime)
	// foramting the Time
	fmt.Println(presentTime.Format("01-05-2006 Monday"))
	// createDate
	createDate:=time.Date(2021,time.August,20,20,12,13,14,time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
    
	// Memory Managment pointer
	myNumber:=50
	var ptr =&myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
	// importent
	*ptr=*ptr+10
	fmt.Println(myNumber)
    


}