package main
import "fmt"
// defer exicute function wish
func Test(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
	fmt.Println("js")
}

func main() {
	defer fmt.Println("golang")
	fmt.Println("defer")
	Test()
	defer fmt.Print("java")
	defer fmt.Print("py")
}

// 0 1 2 3 4 
// golang java py