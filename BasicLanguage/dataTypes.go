package main
import "fmt"

const LoginToken string="pruthiraj@097" // this is public

func main() {
    // %T\n to print the dataType
	// var variableName DataType =value
	var username string="pruthiraj";
	fmt.Println(username)
	fmt.Println("variable type:",username)

	var isLoggedIn bool=true
	fmt.Println("variable type: %T\n",isLoggedIn)
    

	var smallValue uint8 = 255
	fmt.Printf("variable type: %T\n", smallValue)

	var smallFloat float32=245.4747
	fmt.Println(smallFloat)

	//default value and some aliases
	var anotherVariable int
	fmt.Println("variable type: ",anotherVariable)

	//implicite type
	var website="code.go"
	fmt.Println(website)

	// no var style
	// it allowed only inside a function
	numberOfUser:=45
	fmt.Println(numberOfUser)

    fmt.Println(LoginToken)
}