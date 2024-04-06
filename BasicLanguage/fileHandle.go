package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// function to read the file
func ReadFile(fileName string){
	databyte,err:=ioutil.ReadFile(fileName)
	// data are in byte formate
	cheakError(err)
	fmt.Println("Data in file is ",string(databyte))
	// converting to string
}

func main()  {
	fmt.Println("file handaling")
	// creating file using os library
	file,err:=os.Create("./mygolangfile.text")
	cheakError(err)
	// writing in file
	length,err:=io.WriteString(file,"This is my golang file")
	// length is the length of string
	cheakError(err)
		fmt.Println("length is ",length)
	// closing thr file
	file.Close()
	ReadFile("./mygolangfile.text")
	
}

func cheakError(err error)  {
	if err!=nil {
		panic(err)
	}
}