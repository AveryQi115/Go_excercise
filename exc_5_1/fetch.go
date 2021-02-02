package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
)

func main(){
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			log.Fatal("cannot get url")
		}
		body, err := ioutil.ReadAll(resp.Body)

		// it is really important to close response body
		defer resp.Body.Close()
		if err != nil{
			log.Fatal("read response body failed")
		}
		fmt.Printf("%s",body)
	}
}