package main

import(
	"fmt"
	"os"
)

func removeDup(s []string)[]string{
	i := -1
	for _,str := range s{

		if i == -1{
			i++
			continue
		}

		if str != s[i]{
			s[i+1] = str
			i++
		}
	}

	return s[:i+1]
}

func main(){
	if len(os.Args)<=1{
		return
	}

	var strbuf []string
	for _,v := range os.Args[1:]{
		strbuf = append(strbuf, v)
	}

	res := removeDup(strbuf)
	fmt.Println(res)
}