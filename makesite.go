package main

import (
	"os"
	"flag"
	"strings"
	"io/ioutil"
	"html/template"
)

func main() {
	// used $Go build && ./makesite --file=latest-post.txt
	fileName := flag.String("file", "", "the file name to read")
	// used $Go build && ./makesite --dir="./testDir"
	dirName := flag.String("dir", "", "file directory")
	// parses all created flags
	flag.Parse()

	// reads the contents of our file
	fileContents, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
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