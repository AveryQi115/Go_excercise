package main

import "fmt"

type Test struct{
	test_key int
	test_value string
}
func main(){
	//var test Test
	//if test == nil{
	//	fmt.Println("nil type")
	//}
	var test2 []Test
	if test2 == nil{
		fmt.Println("nil slice type")
	}
	var test3 = make([]Test,0)
	if test3 == nil{
		fmt.Println("nil slice type by make")
	}
}
