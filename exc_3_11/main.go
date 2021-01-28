package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main(){
	if len(os.Args)<=1{
		return
	}
	var s string = os.Args[1]
	fmt.Printf("The ans from comma is: %s\n",comma(s))
}

func comma(s string)string{
	var buf bytes.Buffer

	// handle sign
	start := 0
	if s[0]=='+' || s[0]=='-'{
		buf.WriteByte(s[0])
		start = 1
	}

	// handle floating point
	end := strings.Index(s,".")
	if end == -1{
		end = len(s)
	}

	//handle integer part
	pre := (end - start)%3
	if pre == 0{
		pre = 3
	}
	buf.WriteString(s[start:start+pre])

	for i:=start+pre;i+3<=end;i+=3{
		buf.WriteByte(',')
		buf.WriteString(s[i:i+3])
	}

	if end != len(s){
		buf.WriteString(s[end:])
	}

	return buf.String()
}
