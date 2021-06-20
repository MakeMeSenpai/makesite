package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read first post contents
	firstPostContents, _ := ioutil.ReadFile("first-post.txt")

	// write first post contents to new-file.txt
	bytesToWrite := []byte(firstPostContents)
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}

	// read file contents of new-file.txt
	fileContents, err := ioutil.ReadFile("new-file.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
    fmt.Print(string(fileContents))
}
