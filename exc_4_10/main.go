// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func classifier(items *[]*Issue, key string) map[string]*[]Issue {
	if key == "CreatedAt" {
		f := func(item Issue, t time.Time) bool {
			return item.CreatedAt.After(t)
		}

		var monthly []Issue
		var yearly []Issue
		var others []Issue

		for _, item := range *items {
			if f(*item, time.Now().Add(-30*24*time.Hour)) {
				monthly = append(monthly, *item)
			} else if f(*item, time.Now().Add(-365*24*time.Hour)) {
				yearly = append(yearly, *item)
			} else {
				others = append(others, *item)
			}
		}
		return map[string]*[]Issue{"monthly": &monthly, "yearly": &yearly, "others": &others}
	}
	return nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	res := classifier(&result.Items, "CreatedAt")
	if res == nil {
		log.Fatal("no option")
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("*******************Monthly*******************")
	for _, item := range *res["monthly"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("*******************Yearly*******************")
	for _, item := range *res["yearly"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("*******************LongAgo*******************")
	for _, item := range *res["others"] {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}
