package main

import(
	"fmt"
	"os"
	"strconv"
)

// *[size]int means pointer for array
// we need to use pointer because parameters are all passed by value in go
func reverse(s *[5]int){
	fmt.Printf("array pointer type:%T,value:%v\n",s,s)
	tempSlice := s[:]
	fmt.Printf("slice type:%T,value:%v\n",tempSlice,tempSlice)
	for i:=0;i<len(s)/2;i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func main(){
	if len(os.Args)<=1{
		return
	}

	s := [5]int{0,0,0,0,0}
	for i:=0;i<5;i++{
		s[i],_ = strconv.Atoi(os.Args[i+1])
	}
	reverse(&s)
	fmt.Println(s)
}

