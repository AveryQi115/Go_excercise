package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	count := 0
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan(){
		count++
	}
	fmt.Println(count)
}
