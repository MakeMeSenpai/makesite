package main

import (
	"os"
	"log"
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"html/template"
)

// there are two different funcs happening here so look carefully
// fileName reads a files contents then creates a new template
// while dirName reads a directories contents, and changes those into html files
func main() {
	// used $Go build && ./makesite --file=latest-post.txt
	fileName := flag.String("file", "", "the file name to read")
	// used $Go build && ./makesite --dir=.
	dirName := flag.String("dir", "", "file directory to read")
	// parses all created flags
	flag.Parse()

	// reads the contents of our file
	fileContents, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	// reads the files in our directory
    files, err := ioutil.ReadDir(*dirName)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
		// if .txt in file name
		if strings.Contains(file.Name(), ".txt") {
			fmt.Println(file.Name())
			// creates a new file with the name of file.Name found in Directory
			f, err := os.Create(strings.SplitN(file.Name(), ".", 2)[0] + ".html")
			if err != nil {
				panic(err)
			}
		}
    }

	// creates a new file with the name of fileName.html
	f, err := os.Create(strings.SplitN(*fileName, ".", 2)[0] + ".html")
	if err != nil {
		panic(err)
	}

	// parses old template into a new template to pass values
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(f, string(fileContents))
	if err != nil {
		panic(err)
	}
}