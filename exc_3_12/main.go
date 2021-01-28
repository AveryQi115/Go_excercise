package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args)<=2{
		return
	}
	s,t := os.Args[1],os.Args[2]

	ok := compare(s,t)
	fmt.Printf("These 2 strings have the same characters? %v\n",ok)
}

func compare(s,t string)bool{
	sFreq := map[rune]int{}
	for _,v := range s{
		sFreq[v]++
	}
	tFreq := map[rune]int{}
	for _,v := range t{
		tFreq[v]++
	}

	for k,v := range sFreq{
		if v!=tFreq[k]{
			return false
		}
	}
	
	for k,v := range tFreq{
		if v!=sFreq[k]{
			return false
		}
	}
	return true
}
