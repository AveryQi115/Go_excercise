package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main(){
	s,err := ioutil.ReadAll(os.Stdin)
	if err != nil{
		log.Fatal("cannot get standard input")
	}
	var f = func(s string)string{
		return strings.ToUpper(s)
	}
	result := expand(string(s),f)
	fmt.Printf("%s",result)
}

func expand(s string, f func(string)string)string{
	strSlices := strings.Split(s, "foo")
	var buf bytes.Buffer
	for i,strSlice := range strSlices{
		buf.Write([]byte(strSlice))
		if i!=len(strSlices)-1{
			buf.Write([]byte(f("foo")))
		}
	}
	return buf.String()
}
