package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"crypto/sha256"
	"crypto/sha512"
)

var hash = flag.String("hash","sha256","choose hash encoding method")

func main(){
	if len(os.Args)<=1{
		return
	}
	src := os.Args[1]
	switch *hash{
	case "sha256":
		fmt.Println("The hash code for string is:",sha256.Sum256([]byte(src)))
	case "sha384":
		fmt.Println("The hash code for string is:",sha512.Sum512([]byte(src)))
	default:
		log.Fatal(fmt.Sprintf("no option for %s",*hash))
	}
}
