package main

import (
	"bytes"
	"os"
	"fmt"
)

func main(){
	if len(os.Args) <= 1{
		return
	}

	var s string = os.Args[1]
	fmt.Printf("The ans from comma is: %s\n",comma(s))
}



// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	start := len(s)%3
	if start == 0{
		start = 3
	}
	buf.WriteString(s[:start])

	for ;start+3<=len(s);start+=3{
		buf.WriteByte(',')
		buf.WriteString(s[start:start+3])
	}

	return buf.String()
}
