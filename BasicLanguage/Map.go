package main
import "fmt"

func printMap(m map[string]int) int {
	var sum int
	for _,value:=range m{
		sum+=value
	}
	return sum
}

func main()  {
	fmt.Println("maps in golang")
	// declaraing map KEY -VALUE
	language:=make(map[string]string)
	language["js"]="javascript"
	language["py"]="python"
	fmt.Println(language)
	fmt.Println(language["js"])

	//delete 
	delete(language,"py")
    language["java"]="java"
	// looping of map
	// for key,value:=range language{
	// 	fmt.Println(key, value)
	// }


   var hashmap =make(map[string]int)
   hashmap["one"]=1;
   hashmap["two"]=2;
   hashmap["three"]=3;
   fmt.Println(printMap(hashmap))

}