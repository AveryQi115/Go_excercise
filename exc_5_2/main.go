package main

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/net/html"
)

func main(){
	doc,err := html.Parse(os.Stdin)
	if err != nil{
		log.Fatal("Parse Error")
	}
	var records = map[string]int{}
	CountNode(records, doc)
	fmt.Printf("%v",records)
}

func CountNode(records map[string]int, n *html.Node){
	if n.Type == html.ElementNode{
		records[n.Data]++
	}

	if n.FirstChild != nil{
		CountNode(records, n.FirstChild)
	}

	if n.NextSibling != nil{
		CountNode(records, n.NextSibling)
	}
}