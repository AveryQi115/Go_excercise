package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// search issues
func search(args []string) {
	res, err := SearchIssues(args)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to search issues %s", err.Error()))
	}
	items := res.Items
	fmt.Printf("\t%-10s\t%-20s\n", "User", "Title")
	for _, item := range items {
		fmt.Printf("\t%-10s\t%-20s\n", item.User, item.Title)
	}
}

func read(owner string, repo string, number string) {
	item, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to get issue:%s", err.Error()))
	}
	fmt.Printf("User:\t%-10s,\tTitle:\t%-20s\n", item.User, item.Title)
	fmt.Print(item.Body)
}

func close_(owner string, repo string, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{
		"state": "closed",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func open(owner string, repo string, number string) {
	_, err := EditIssue(owner, repo, number, map[string]string{
		"state": "open",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func edit(owner string, repo string, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	//
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}

	// create a tempFile
	// eg:in vim it's a swp file
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	// change the searched issue into json format
	// write the json file to tempFile
	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	// editorPath vim filename
	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// move the file pointer to where the file starts
	tempfile.Seek(0, 0)
	fields := make(map[string]string)

	// decode what user input to json format
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	// editIssue across http patch method
	_, err = EditIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

// prompt usage for invalid command
// use ` to include operative characters like \n and \t etc
var usage string = `usage:
search QUERY
[read|edit|close|open] OWNER REPO ISSUE_NUMBER`

// error caused by invalid usage
func usageDie() {
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageDie()
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	if cmd == "search" {
		if len(args) < 1 {
			usageDie()
		}
		search(args)

		//TODO: why there's not return ?
		os.Exit(0)
	}
	if len(args) != 3 {
		usageDie()
	}
	owner, repo, number := args[0], args[1], args[2]
	switch cmd {
	case "read":
		read(owner, repo, number)
	case "edit":
		edit(owner, repo, number)
	case "close":
		close_(owner, repo, number)
	case "open":
		open(owner, repo, number)
	}
}
