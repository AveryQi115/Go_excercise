package main

import (
	"fmt"
	"os"
)

func main(){
	nums := os.Args[1:]
	//TODO: change str to int
	fmt.Printf("%v\n", max(nums[0], nums[1:]))
	fmt.Printf("%v\n", min(nums[0], nums[1:]))
}

func max(first int, vals...int)int{
	max := first
	for _,val := range vals{
		if val>max{
			max = val
		}
	}
	return max
}

func min(first int,vals...int)int{
	min := first
	for _,val := range vals{
		if val <min{
			min = val
		}
	}
	return min
}
