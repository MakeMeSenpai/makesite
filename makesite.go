package main

import (
	"os"
	"flag"
	"strings"
	"io/ioutil"
	"html/template"
)

func main() {
	// flag
	fileName := flag.String("file", "", "the file name to read")
	flag.Parse()

	fileContents, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(strings.SplitN(*fileName, ".", 2)[0] + ".html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(f, string(fileContents))
	if err != nil {
		panic(err)
	}
}
